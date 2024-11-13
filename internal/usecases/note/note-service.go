package note

import (
	"ListBotTG/internal/usecases"
)

type Service struct {
	repo usecases.INote
}

func NewNoteService(repo usecases.INote) *Service {
	return &Service{repo: repo}
}

var _ usecases.INoteManagementSystem = (*Service)(nil)
