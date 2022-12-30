package handlers

import (
	"github.com/TulioGuaraldoB/cross-roads/core/controllers"
	"github.com/gin-gonic/gin"
)

func NewUserHandler(router *gin.RouterGroup, userController controllers.IUserController) {
	user := router.Group("user")
	{
		user.POST("login", userController.Login)
	}
}
