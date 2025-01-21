package app

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
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

func NewApp(Env *config.EnvironmentVariables) *App {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            *aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewEnvCredentials()),
	}))

	// Create DynamoDB client
	dynamo := dynamodb.New(sess)

	userRepository := repositories.NewUserRepository(dynamo)
	userService := services.NewUserService(userRepository)

	mux := chi.NewMux()
	mux.Use(middleware.Logger)
	mux.Use(middleware.SetHeader("Content-Type", "application/json"))
	userService.RegisterHandlers(mux)

	return &App{
		Env:     Env,
		Handler: mux,
	}
}
