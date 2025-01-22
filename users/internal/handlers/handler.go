package handlers

import "github.com/go-chi/chi/v5"

type AppHandler interface {
	RegisterHandlers(*chi.Mux)
}

const (
	CreateUserErrorCode int8 = iota
)

const (
	UserServiceName = "user_service"
)

type ErrorResponse struct {
	Code        int8   `json:"code"`
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	Operation   string `json:"operation"`
}
