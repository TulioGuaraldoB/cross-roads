package controllers

import (
	"net/http"

	"github.com/TulioGuaraldoB/cross-roads/core/businesses"
	"github.com/TulioGuaraldoB/cross-roads/core/dtos/requests"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type userController struct {
	userBusiness businesses.IUserBusiness
}

func NewUserController(userBusiness businesses.IUserBusiness) IUserController {
	return &userController{
		userBusiness: userBusiness,
	}
}

func (c *userController) Login(ctx *gin.Context) {
	userCredentials := new(requests.Credentials)
	if err := ctx.ShouldBindJSON(userCredentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	token, err := c.userBusiness.LoginUser(ctx, userCredentials)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, *token)
}

func (c *userController) Register(ctx *gin.Context) {
	userReq := new(requests.UserRequest)
	if err := ctx.ShouldBindJSON(userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	_, err := c.userBusiness.RegisterUser(ctx, userReq)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user inserted successfully!",
	})
}
