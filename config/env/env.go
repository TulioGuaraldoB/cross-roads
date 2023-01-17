package env

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVar struct {
	KeyCloakBasePath      string
	KeyCloakRealm         string
	KeyCloakClientId      string
	KeyCloakClientSecret  string
	KeyCloakAdminUsername string
	KeyCloakAdminPassword string
	Port                  string
}

var Env *EnvironmentVar

func GetEnvironmentVariables() *EnvironmentVar {
	godotenv.Load(".env")
	Env = &EnvironmentVar{
		KeyCloakBasePath:      os.Getenv("KEYCLOAK_BASE_PATH"),
		KeyCloakRealm:         os.Getenv("KEYCLOAK_REALM"),
		KeyCloakClientId:      os.Getenv("KEYCLOAK_CLIENT_ID"),
		KeyCloakClientSecret:  os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		KeyCloakAdminUsername: os.Getenv("KEYCLOAK_ADMIN_USERNAME"),
		KeyCloakAdminPassword: os.Getenv("KEYCLOAK_ADMIN_PASSWORD"),
		Port:                  os.Getenv("PORT"),
	}

	return Env
}
