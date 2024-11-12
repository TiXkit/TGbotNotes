package note

import (
	"ListBotTG/internal/usecases"
	"context"
)

func (s *Service) DropNote(ctx context.Context, note usecases.NoteDTO) (bool, error) {
	err := s.repo.DropNote(ctx, &note.Note)
	if err != nil {
		return false, err
	}
	return true, nil
}
