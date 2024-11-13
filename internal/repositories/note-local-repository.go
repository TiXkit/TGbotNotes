package repositories

import (
	"ListBotTG/internal/configs"
	"ListBotTG/internal/models"
	"ListBotTG/internal/usecases"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type NoteLocalRepository struct {
	PathToStorage string
}

func NewNoteRepos(db *sql.DB) *NoteLocalRepository {
	return &NoteLocalRepository{}
}

var _ usecases.INote = (*NoteLocalRepository)(nil)

func (nr *NoteLocalRepository) AddNote(ctx context.Context, note *models.Note) (int, error) {
	file, err := os.OpenFile(nr.PathToStorage, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return 0, fmt.Errorf("ошибка при открытии файла")
	}
	defer file.Close()

	//jsonStructForNoteDecoder := struct {
	//ID int
	//Name string}{} или лучше так?
	var notesDecoder usecases.NoteDTO // Стоит ли так делать или нужно объявлять для этого новую структуру
	counterIDJson := 0

	for {
		if err := json.NewDecoder(file).Decode(&notesDecoder); err != nil {
			if err.Error() == "EOF" {
				break // конец файла
			}
			return 0, fmt.Errorf("ошибка при декодировании файла")
		}

		if notesDecoder.Email == note.Email {
			return 0, fmt.Errorf("данный email уже существует в базе")
		}
		counterIDJson++ // высчитываем последний ID
	}

	note.ID = models.IDNote(counterIDJson + 1) // присваиваем полю переменной типа структуры ID
	noteInByte, err := json.Marshal(note)
	if err != nil {
		return 0, err
	}
	if _, err := file.WriteString(string(noteInByte) + "\n"); err != nil { // добавляем в файл
		return 0, fmt.Errorf("ошибка при записи в файл")
	}

	return counterIDJson + 1, nil
}

func (nr *NoteLocalRepository) DropNote(ctx context.Context, note *models.Note) error {
	file, err := os.OpenFile(nr.PathToStorage, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при открытии файла")
	}
	defer file.Close()

	var searchField interface{}
	if note.Email != "" {
		searchField = note.Email
	} else {
		searchField = note.ID
	}
	var counterIDJson int
	var notesDecoder usecases.NoteDTO
	fileFound := false

	for {
		if err := json.NewDecoder(file).Decode(&notesDecoder); err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("ошибка при декодировании файла")
		}
		if notesDecoder.Email == searchField || notesDecoder.ID == searchField {
			fileFound = true // Если мы нашли мыло
			break
		}
		counterIDJson++
	}

	if !fileFound {
		return fmt.Errorf("данного email нет в списке")
	}

	fileReadBytes, err := os.ReadFile(configs.PathFileLocalStorage)
	if err != nil {
		return fmt.Errorf("ошибка при чтении файла")
	}
	linesFile := strings.Split(string(fileReadBytes), "\n") // разбирает строку на массив, с помощью разделителя
	if counterIDJson >= 0 && counterIDJson < len(linesFile) {
		linesFile = append(linesFile[:counterIDJson], linesFile[counterIDJson+1:]...)
	}
	newData := strings.Join(linesFile, "\n") // собирает строку из массива, добавляя разделитель

	if err := os.WriteFile(configs.PathFileLocalStorage, []byte(newData), os.ModePerm); err != nil {
		return fmt.Errorf("ошибка при перезаписи файла")
	}

	return nil
}

func (nr *NoteLocalRepository) ShowListNote(ctx context.Context, page1, page2 int) ([]usecases.NoteDTO, error) {
	file, err := os.OpenFile(nr.PathToStorage, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла")
	}
	defer file.Close()

	// Нужно от n-page1 показать все записи до n-page2, (2 <= (Npage2 - Npage1) <= 100)
	var notesDecoder usecases.NoteDTO
	sliceNotesForDecoding := make([]usecases.NoteDTO, page2-page1, page2-page1)

	for i := 0; i < (page2 + 1); {
		if err := json.NewDecoder(file).Decode(&notesDecoder); err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("ошибка при декодировании файла")
		}
		if (i >= page1) && (i <= page2) {
			sliceNotesForDecoding[i] = notesDecoder
		}
		i++
	}
	return sliceNotesForDecoding, nil
}

// доделать последний алгоритм(учесть то, что введенного интервала в файле может не быть)
// и просмотреть все предыдущие алгоритмы на работоспособность.
