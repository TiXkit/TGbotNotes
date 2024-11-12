package usecases

import (
	"ListBotTG/internal/models"
	"context"
)

type INote interface {
	AddNote(ctx context.Context, note *models.Note) (int, error)
	DropNote(ctx context.Context, note *models.Note) error
	ShowListNote(ctx context.Context, page1, page2 int) ([]models.Note, error)
}
