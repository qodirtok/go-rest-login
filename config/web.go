package config

import (
	"net/http"

	"githiub.com/qodirtok/go-rest-login/controller"
	"github.com/gin-gonic/gin"
)

func Web(router *gin.Engine)  {

	router.GET("ping",func(ctx *gin.Context) {

		 ctx.AsciiJSON(http.StatusOK,gin.H{"message":"pong"})
	})

	r := router.Group("/auth")

	r.POST("/register",controller.Register)
	r.POST("/login",controller.Login)
}