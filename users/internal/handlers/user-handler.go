package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nimbo1999/lopes_drive/users/internal/handlers/dto"
	"github.com/nimbo1999/lopes_drive/users/internal/models"
	"github.com/nimbo1999/lopes_drive/users/internal/repositories"
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
		r.Get("/{username}", handler.getUser)
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
		if err == services.ErrRecordAlreadyExist {
			response.WriteHeader(http.StatusConflict)
			encoder.Encode(ErrorResponse{
				Code:        CreateUserAlreadyExist,
				Message:     err.Error(),
				ServiceName: UserServiceName,
				Operation:   "create",
			})
			return
		}

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

func (handler *userHandler) getUser(response http.ResponseWriter, request *http.Request) {
	username := chi.URLParam(request, "username")
	encoder := json.NewEncoder(response)
	if strings.TrimSpace(username) == "" {
		log.Println("error:", ErrPathParameterInvalid.Error())
		response.WriteHeader(http.StatusBadRequest)
		encoder.Encode(ErrorResponse{
			Code:        PathParameterErrorCode,
			Message:     ErrPathParameterInvalid.Error(),
			ServiceName: UserServiceName,
			Operation:   "create",
		})
		return
	}

	user, err := handler.UserService.FindById(username)
	if err != nil {
		log.Println("error:", err.Error())
		if err == repositories.ErrRecordNotFound {
			response.WriteHeader(http.StatusNotFound)
			encoder.Encode(ErrorResponse{
				Code:        GetUserByUsernameNotFoundErrorCode,
				Message:     err.Error(),
				ServiceName: UserServiceName,
				Operation:   "get_by_username",
			})
			return
		}

		response.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ErrorResponse{
			Code:        GetUserByUsernameErrorCode,
			Message:     err.Error(),
			ServiceName: UserServiceName,
			Operation:   "get_by_username",
		})
		return
	}

	response.WriteHeader(http.StatusOK)
	encoder.Encode(dto.CreateUserResponse{
		Username: user.Username,
		Email:    user.Email,
	})
}
