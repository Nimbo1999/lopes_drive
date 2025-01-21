package repositories

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/nimbo1999/lopes_drive/users/internal/models"
)

var ErrRepositoryMethodNotImplementedYet = errors.New("repository: method not implemented yet")

type userRepository struct {
	dynamo *dynamodb.DynamoDB
}

func NewUserRepository(db *dynamodb.DynamoDB) *userRepository {
	return &userRepository{dynamo: db}
}

func (repository *userRepository) FindById(id string) (*models.User, error) {
	return nil, ErrRepositoryMethodNotImplementedYet
}

func (repository *userRepository) Delete(user models.User) error {
	return ErrRepositoryMethodNotImplementedYet
}

func (repository *userRepository) Update(user models.User) error {
	return ErrRepositoryMethodNotImplementedYet
}

func (repository *userRepository) List() ([]models.User, error) {
	return nil, ErrRepositoryMethodNotImplementedYet
}

// @TODO: Create the Repository Methods
