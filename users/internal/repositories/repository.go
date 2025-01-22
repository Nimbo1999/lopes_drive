package repositories

import "errors"

type Repository[Model any] interface {
	Create(*Model) error
	FindById(id string) (*Model, error)
	Delete(any) error
	Update(any) error
	List() ([]Model, error)
}

var (
	ErrRepositoryMethodNotImplementedYet = errors.New("repository: method not implemented yet")
)
