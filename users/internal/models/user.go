package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(username, email string) (*User, error) {
	return &User{
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
