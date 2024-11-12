package note

import (
	"context"
	"fmt"
)

func (s *Service) ShowListNote(ctx context.Context, page1, page2 int) error {
	if page1 >= 2 && (page2-page1) <= 100 {
		s.repo.ShowListNote(ctx, page1, page2)
	} else {
		return fmt.Errorf("слишком большое значение поиска")
	}
	return nil
}
