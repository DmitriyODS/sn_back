package app

import (
	"idon.com/models"
)

type App struct {
	reps models.Repository
}

func MakeApp(reps models.Repository) *App {
	return &App{reps: reps}
}
