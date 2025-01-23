package services

import (
	"github.com/nimbo1999/lopes_drive/users/internal/handlers/dto"
	"github.com/nimbo1999/lopes_drive/users/internal/models"
	"github.com/nimbo1999/lopes_drive/users/internal/repositories"
)

type userService struct {
	repository repositories.Repository[models.User]
}

func NewUserService(repository repositories.Repository[models.User]) *userService {
	return &userService{repository}
}

func (service *userService) Create(p any) (*models.User, error) {
	payload, ok := p.(dto.CreateUserPayload)
	if !ok {
		return nil, ErrInvalidPayload
	}

	user, err := models.NewUser(payload.Username, payload.Email)
	if err != nil {
		return nil, err
	}

	if _, err = service.repository.FindById(user.Username); err != repositories.ErrRecordNotFound {
		return nil, ErrRecordAlreadyExist
	}

	if err = user.SetPassword(payload.Password); err != nil {
		return nil, err
	}

	if err = service.repository.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (service *userService) FindById(id string) (*models.User, error) {
	return service.repository.FindById(id)
}
