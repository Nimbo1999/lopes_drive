package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Username string
	Email    string
	Password string
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

func (user *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	return nil
}

func (user User) ComparePassword(storegedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(storegedPassword))
}
