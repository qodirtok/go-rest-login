package cmd

import (
	"githiub.com/qodirtok/go-rest-login/config"
	"github.com/gin-gonic/gin"
)


func Cmd()  {
	config.Connect()

	r := gin.Default()

	config.Web(r)

	r.Run(":4747")
}