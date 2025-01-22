package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/nimbo1999/lopes_drive/users/internal/models"
)

type userRepository struct {
	dynamo *dynamodb.DynamoDB
}

func NewUserRepository(db *dynamodb.DynamoDB) *userRepository {
	return &userRepository{db}
}

func (repository *userRepository) Create(model *models.User) error {
	item, err := dynamodbattribute.MarshalMap(model)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("lopes-drive_users"),
		Item:      item,
	}

	if _, err := repository.dynamo.PutItem(input); err != nil {
		return err
	}
	return nil
}

func (repository *userRepository) FindById(id string) (*models.User, error) {
	return nil, ErrRepositoryMethodNotImplementedYet
}

func (repository *userRepository) Delete(payload any) error {
	return ErrRepositoryMethodNotImplementedYet
}

func (repository *userRepository) Update(payload any) error {
	return ErrRepositoryMethodNotImplementedYet
}

func (repository *userRepository) List() ([]models.User, error) {
	return nil, ErrRepositoryMethodNotImplementedYet
}
