package controller

import (
	"net/http"

	"githiub.com/qodirtok/go-rest-login/helper"
	"githiub.com/qodirtok/go-rest-login/services"
	"github.com/gin-gonic/gin"
)


func Login(context *gin.Context)  {
	var input services.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest , gin.H{"err":err.Error()})
		return
	}

	user, err := services.FindUserByUsername(input.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	}


	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(services.Users{})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}