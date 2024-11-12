package repositories

import (
	"ListBotTG/internal/models"
	"database/sql"
)

type NoteDBRepository struct {
	db *sql.DB
}

func NewNoteDBRepository(db *sql.DB) *NoteDBRepository {
	return &NoteDBRepository{db: db}
}

func (nr *NoteDBRepository) AddNote(note *models.Note) (int, error) {
	return 0, nil
}
func (nr *NoteDBRepository) DropNote(note *models.Note) error {
	return nil
}
func (nr *NoteDBRepository) ShowListNote(page1, page2 int) ([]models.Note, error) {

	return nil, nil
}
