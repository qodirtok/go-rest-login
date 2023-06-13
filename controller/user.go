package controller

import (
	"net/http"

	"githiub.com/qodirtok/go-rest-login/services"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context)  {
	var input services.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"msg" : err.Error()})
		return
	}

	user := services.Users{
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"msg" : err.Error()})
		return
	}
	context.AsciiJSON(http.StatusOK,savedUser)
}