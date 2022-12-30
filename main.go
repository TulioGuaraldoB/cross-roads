package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Nerzal/gocloak/v12"
	"github.com/TulioGuaraldoB/cross-roads/config/env"
	"github.com/TulioGuaraldoB/cross-roads/infra/server"
)

func main() {
	env.GetEnvironmentVariables()

	server := server.New()
	server.Run()

	// cloakClient := gocloak.NewClient("http://localhost:8082")
	// ctx := context.Background()

	// loginUser(ctx, cloakClient)
	// token, err := cloakClient.LoginAdmin(ctx, "admin", "admin", "master")
	// if err != nil {
	// 	errMessage := fmt.Sprintf("Failed to login admin in KeyCloak. %s", err.Error())
	// 	log.Fatal(errMessage)
	// }

	// createUser(ctx, cloakClient, token)
}

func createUser(ctx context.Context, keyCloakClient *gocloak.GoCloak, jwtCloak *gocloak.JWT) {
	newUser := gocloak.User{
		FirstName: gocloak.StringP("Bilbo"),
		LastName:  gocloak.StringP("Baggings"),
		Email:     gocloak.StringP("bilbo.baggings@hotmail.com"),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP("baggingB"),
	}

	userID, err := keyCloakClient.CreateUser(ctx, jwtCloak.AccessToken, "master", newUser)
	if err != nil {
		errMessage := fmt.Sprintf("Failed to register user. %s", err.Error())
		log.Fatal(errMessage)
	}

	fmt.Printf("created user ID: %s", userID)
}

func loginUser(ctx context.Context, keyCloakClient *gocloak.GoCloak) {
	token, err := keyCloakClient.Login(ctx, "car-app", "Wr8uPTOsKwQXky1L6ikekD3SCURCnEUG", "master", "baggingB", "bags")
	if err != nil {
		errMessage := fmt.Sprintf("Failed to login. %s", err.Error())
		log.Fatal(errMessage)
	}

	fmt.Println(fmt.Sprintf("Success Login! Token: %s", token.AccessToken))
}
