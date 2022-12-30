package env

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVar struct {
	KeyCloakBasePath     string
	KeyCloakRealm        string
	KeyCloakClientId     string
	KeyCloakClientSecret string
	Port                 string
}

var Env *EnvironmentVar

func GetEnvironmentVariables() *EnvironmentVar {
	godotenv.Load(".env")
	Env = &EnvironmentVar{
		KeyCloakBasePath:     os.Getenv("KEYCLOAK_BASE_PATH"),
		KeyCloakRealm:        os.Getenv("KEYCLOAK_REALM"),
		KeyCloakClientId:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		KeyCloakClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		Port:                 os.Getenv("PORT"),
	}

	return Env
}
