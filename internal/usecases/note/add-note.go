package note

import (
	"ListBotTG/internal/models"
	"ListBotTG/internal/usecases"
	"context"
	"errors"
)

func (s *Service) AddNote(ctx context.Context, note usecases.NoteDTO) (int, error) {
	addNote := models.Note{
		ID:    models.IDNote(note.ID),
		Email: note.Email,
	}

	id, err := s.repo.AddNote(ctx, &addNote)
	if err != nil {
		if errors.Is(err, models.InvalidEmail) {
			return 0, models.InvalidEmail
		} else {
			return 0, models.ErrorFromLocalStorage
		}
	}

	return id, nil
}
