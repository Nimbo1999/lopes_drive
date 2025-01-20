package models

import "github.com/google/uuid"

type User struct {
	ID       string
	Username string
	Email    string
}

func NewUser(username, email string) (*User, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       uuid.String(),
		Username: username,
		Email:    email,
	}, nil
}
