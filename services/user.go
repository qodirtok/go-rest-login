package services

import (
	"html"
	"strings"

	"githiub.com/qodirtok/go-rest-login/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users models.User

func (user *Users) Save() (*Users,error)  {
	err := models.Database.Create(&user).Error

	if err != nil {
		return &Users{},err
	}

	return user,nil
}

func (user *Users) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}

func FindUserByUsername(username string) (Users, error) {
	var users Users
	err := models.Database.Where("username=?",username).Find(&users).Error
	if err != nil {
		return Users{},nil
	}
	return users,nil
}