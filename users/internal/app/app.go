package app

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nimbo1999/lopes_drive/users/internal/config"
	"github.com/nimbo1999/lopes_drive/users/internal/repositories"
	"github.com/nimbo1999/lopes_drive/users/internal/services"
)

type App struct {
	Handler http.Handler
	Env     *config.EnvironmentVariables
}

type AppParams struct {
	Db  *dynamodb.DynamoDB
	Env *config.EnvironmentVariables
}

func NewApp(params *AppParams) *App {
	userRepository := repositories.NewUserRepository(params.Db)
	userService := services.NewUserService(userRepository)

	mux := chi.NewMux()
	mux.Use(middleware.Logger)
	mux.Use(middleware.SetHeader("Content-Type", "application/json"))
	userService.RegisterHandlers(mux)

	return &App{
		Env:     params.Env,
		Handler: mux,
	}
}
