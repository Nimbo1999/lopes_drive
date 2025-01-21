package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nimbo1999/lopes_drive/users/internal/models"
	"github.com/nimbo1999/lopes_drive/users/internal/repositories"
)

type userService struct {
	repository repositories.Repository[models.User]
}

func NewUserService(repository repositories.Repository[models.User]) *userService {
	return &userService{repository}
}

// @ TODO: Create the service Methods
func (service *userService) RegisterHandlers(mux *chi.Mux) {
	mux.Get("/", service.ListUsers)
}

func (service *userService) ListUsers(response http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(response)
	matheus, _ := models.NewUser("matheus", "matlopes1999@gmail.com")
	suzana, _ := models.NewUser("suzana", "suzana@gmail.com")

	mock := GetUserListResponse{
		Data: []*models.User{
			matheus,
			suzana,
		},
		Pagination: Pagination{
			Page:  "1",
			Total: "2",
		},
	}
	if err := encoder.Encode(mock); err != nil {
		log.Fatal(err)
	}
}
