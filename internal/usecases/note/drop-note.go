package note

import (
	"ListBotTG/internal/models"
	"ListBotTG/internal/usecases"
	"context"
	"errors"
)

func (s *Service) DropNote(ctx context.Context, note usecases.NoteDTO) (bool, error) {
	delNote := models.Note{
		ID:    models.IDNote(note.ID),
		Email: note.Email,
	}

	err := s.repo.DropNote(ctx, &delNote)
	if err != nil {
		if errors.Is(err, models.EmailNotFound) {
			return false, models.EmailNotFound
		} else {
			errors.Is(err, models.ErrorFromLocalStorage)
		}
	}

	return true, nil
}
