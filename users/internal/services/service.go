package services

import (
	"errors"

	"github.com/nimbo1999/lopes_drive/users/internal/models"
)

var (
	ErrInvalidPayload = errors.New("repository: invalid payload type")
)

type Service[Model any] interface {
	Create(payload any) (*Model, error)
	FindById(id string) (*Model, error)
}

type Pagination struct {
	Page  string `json:"page"`
	Total string `json:"total"`
}

type GetUserListResponse struct {
	Data       []*models.User `json:"data"`
	Pagination Pagination     `json:"pagination"`
}
