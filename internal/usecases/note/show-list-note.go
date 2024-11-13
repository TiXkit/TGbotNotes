package note

import (
	"ListBotTG/internal/models"
	"ListBotTG/internal/usecases"
	"context"
)

func (s *Service) ShowListNote(ctx context.Context, page1, page2 int) ([]usecases.NoteDTO, error) {

	if !(page1 >= 2 && (page2-page1) <= 100) {
		return nil, models.RangeOverflow
	}

	answerSliceNote, err := s.repo.ShowListNote(ctx, page1, page2)
	if err != nil {
		return nil, models.ErrorFromLocalStorage
	}

	return answerSliceNote, nil
}
