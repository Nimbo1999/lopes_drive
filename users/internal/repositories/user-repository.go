package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/nimbo1999/lopes_drive/users/internal/models"
)

type userRepository struct {
	dynamo    *dynamodb.DynamoDB
	tableName *string
}

func NewUserRepository(db *dynamodb.DynamoDB) *userRepository {
	return &userRepository{dynamo: db, tableName: aws.String("lopes-drive_users")}
}

func (repository *userRepository) Create(model *models.User) error {
	item, err := dynamodbattribute.MarshalMap(model)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: repository.tableName,
		Item:      item,
	}

	if _, err := repository.dynamo.PutItem(input); err != nil {
		return err
	}
	return nil
}

func (repository *userRepository) FindById(id string) (*models.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: repository.tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(id)},
		},
	}

	output, err := repository.dynamo.GetItem(input)
	if err != nil {
		return nil, err
	}

	if output.Item == nil {
		return nil, ErrRecordNotFound
	}

	var user models.User
	if err = dynamodbattribute.UnmarshalMap(output.Item, &user); err != nil {
		return nil, err
	}
	return &user, nil
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
