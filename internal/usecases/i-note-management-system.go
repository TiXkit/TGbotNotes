package usecases

import "context"

type INoteManagementSystem interface {
	AddNote(ctx context.Context, note NoteDTO) (int, error)
	DropNote(ctx context.Context, note NoteDTO) (bool, error)
	ShowListNote(ctx context.Context, page1, page2 int) ([]NoteDTO, error)
}
