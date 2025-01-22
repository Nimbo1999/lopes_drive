package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nimbo1999/lopes_drive/users/internal/handlers/dto"
	"github.com/nimbo1999/lopes_drive/users/internal/models"
	"github.com/nimbo1999/lopes_drive/users/internal/services"
)

type userHandler struct {
	UserService services.Service[models.User]
}

func NewUserHandler(service services.Service[models.User]) *userHandler {
	return &userHandler{service}
}

func (handler *userHandler) RegisterHandlers(mux *chi.Mux) {
	mux.Route("/user", func(r chi.Router) {
		r.Use(middleware.StripSlashes)
		r.Post("/", handler.createUser)
	})
}

func (handler *userHandler) createUser(response http.ResponseWriter, request *http.Request) {
	var payload dto.CreateUserPayload
	encoder := json.NewEncoder(response)
	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		log.Println("error:", err.Error())
		response.WriteHeader(http.StatusBadRequest)
		encoder.Encode(ErrorResponse{
			Code:        CreateUserErrorCode,
			Message:     err.Error(),
			ServiceName: UserServiceName,
			Operation:   "create",
		})
		return
	}

	user, err := handler.UserService.Create(payload)
	if err != nil {
		log.Println("error:", err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ErrorResponse{
			Code:        CreateUserErrorCode,
			Message:     err.Error(),
			ServiceName: UserServiceName,
			Operation:   "create",
		})
		return
	}

	response.WriteHeader(http.StatusCreated)
	encoder.Encode(dto.CreateUserResponse{
		Username: user.Username,
		Email:    user.Email,
	})
}
