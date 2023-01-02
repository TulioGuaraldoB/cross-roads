package server

import (
	"log"

	"github.com/Nerzal/gocloak/v12"
	"github.com/TulioGuaraldoB/cross-roads/config/env"
	"github.com/TulioGuaraldoB/cross-roads/core/businesses"
	"github.com/TulioGuaraldoB/cross-roads/core/controllers"
	"github.com/TulioGuaraldoB/cross-roads/core/handlers"
	"github.com/TulioGuaraldoB/cross-roads/infra/server/routes"
	"github.com/TulioGuaraldoB/cross-roads/infra/services"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Server *gin.Engine
}

func New() Server {
	return Server{
		Port:   env.Env.Port,
		Server: gin.Default(),
	}
}

func (s *Server) Run() {
	// Services
	keyCloakClient := gocloak.NewClient(env.Env.KeyCloakBasePath)
	routerEngine, router := routes.SetRoutes()

	keyCloakService := services.NewKeyCloakService(keyCloakClient)

	// Businesses
	userBusiness := businesses.NewUserBusinesses(keyCloakService)

	// Controllers
	userController := controllers.NewUserController(userBusiness)

	// Handlers
	handlers.NewUserHandler(router, userController)

	log.Fatal(routerEngine.Run(":" + s.Port))
}
