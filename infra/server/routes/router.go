package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	apiVersion := new(gin.RouterGroup)

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		v1.GET("health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "Healthy version!")
		})

		apiVersion = v1
	}

	return router, apiVersion
}
