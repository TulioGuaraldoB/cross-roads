package services

import (
	"context"

	"github.com/Nerzal/gocloak/v12"
	"github.com/TulioGuaraldoB/cross-roads/config/env"
	"github.com/TulioGuaraldoB/cross-roads/core/dtos/requests"
)

type IKeyCloakService interface {
	Login(ctx context.Context, credentials *requests.Credentials) (*string, error)
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
