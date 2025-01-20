package app

import (
	"net/http"
)

type App struct {
	Handler http.Handler
}

func NewApp() *App {
	return &App{}
}
