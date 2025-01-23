package handlers

import (
	"errors"

	"github.com/go-chi/chi/v5"
)

type AppHandler interface {
	RegisterHandlers(*chi.Mux)
}

// Generic Handler error codes
const (
	PathParameterErrorCode int16 = iota + 1
)

// User error codes
const (
	CreateUserErrorCode int16 = iota + 1000
	// This is a generic error code that happens when calling the get user by username handler
	GetUserByUsernameErrorCode
	// This is error code that showup when the user was not found by the username
	GetUserByUsernameNotFoundErrorCode
)

const (
	UserServiceName = "user_service"
)

var (
	ErrPathParameterInvalid = errors.New("handler: the path parameter is invalid")
)

type ErrorResponse struct {
	Code        int16  `json:"code"`
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	Operation   string `json:"operation"`
}
