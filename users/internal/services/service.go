package services

import (
	"github.com/go-chi/chi/v5"
	"github.com/nimbo1999/lopes_drive/users/internal/models"
)

type Service interface {
	RegisterHandlers(*chi.Mux)
}

type Pagination struct {
	Page  string `json:"page"`
	Total string `json:"total"`
}

type GetUserListResponse struct {
	Data       []*models.User `json:"data"`
	Pagination Pagination     `json:"pagination"`
}
