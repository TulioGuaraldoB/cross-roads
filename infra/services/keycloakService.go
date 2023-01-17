package services

import (
	"context"

	"github.com/Nerzal/gocloak/v12"
	"github.com/TulioGuaraldoB/cross-roads/config/env"
	"github.com/TulioGuaraldoB/cross-roads/core/dtos/requests"
)

type IKeyCloakService interface {
	Login(ctx context.Context, credentials *requests.Credentials) (*string, error)
	CreateUser(ctx context.Context, user *requests.UserRequest) (*string, error)
}

type keyCloakService struct {
	keyCloakClient *gocloak.GoCloak
}

func NewKeyCloakService(keyCloakClient *gocloak.GoCloak) IKeyCloakService {
	return &keyCloakService{
		keyCloakClient: keyCloakClient,
	}
}

func (s *keyCloakService) Login(ctx context.Context, credentials *requests.Credentials) (*string, error) {
	token, err := s.keyCloakClient.Login(ctx, env.Env.KeyCloakClientId, env.Env.KeyCloakClientSecret,
		env.Env.KeyCloakRealm, credentials.Username, credentials.Password)
	if err != nil {
		return nil, err
	}

	return &token.AccessToken, nil
}

func (s *keyCloakService) CreateUser(ctx context.Context, user *requests.UserRequest) (*string, error) {
	keyCloakUser := gocloak.User{
		FirstName: gocloak.StringP(user.FirstName),
		LastName:  gocloak.StringP(user.LastName),
		Email:     gocloak.StringP(user.Email),
		Username:  gocloak.StringP(user.Username),
		Credentials: &[]gocloak.CredentialRepresentation{
			{
				SecretData: gocloak.StringP(user.Password),
			},
		},
	}

	token, err := s.keyCloakClient.LoginAdmin(ctx, env.Env.KeyCloakAdminUsername, env.Env.KeyCloakAdminUsername, env.Env.KeyCloakRealm)
	if err != nil {
		return nil, err
	}

	creationToken, err := s.keyCloakClient.CreateUser(ctx, token.AccessToken, env.Env.KeyCloakRealm, keyCloakUser)
	if err != nil {
		return nil, err
	}

	return &creationToken, nil
}
