package note

import (
	"ListBotTG/internal/usecases"
	"context"
)

func (s *Service) AddNote(ctx context.Context, note usecases.NoteDTO) (int, error) {
	id, err := s.repo.AddNote(ctx, &note.Note)
	if err != nil {
		return 0, nil
	}
	return id, nil
}
