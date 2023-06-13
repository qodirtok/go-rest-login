package services

import (
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *Users) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
}