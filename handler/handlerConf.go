package handler

import (
	"github.com/DulatKuntu/bilim/controller"
)

type AppHandler struct {
	Repo *controller.DatabaseRepository
}

func NewHandler(Repo *controller.DatabaseRepository) *AppHandler {
	return &AppHandler{Repo: Repo}
}
