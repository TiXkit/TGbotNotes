package controllers

import "ListBotTG/internal/usecases"

type UseCases struct {
	usecases.INoteManagementSystem
}

type Controller struct {
	UseCases
}

func NewConroller(us UseCases) *Controller {
	return &Controller{UseCases{us}}
}
