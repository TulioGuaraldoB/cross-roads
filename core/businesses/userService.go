package businesses

import (
	"context"

	"github.com/Nerzal/gocloak/v12"
	"github.com/TulioGuaraldoB/cross-roads/config/env"
	"github.com/TulioGuaraldoB/cross-roads/core/dtos/requests"
)

type IUserBusiness interface {
	LoginUser(ctx context.Context, credentials *requests.Credentials) (*string, error)
}

type userBusiness struct {
	keyCloakClient *gocloak.GoCloak
}

func NewUserBusinesses(keyCloakClient *gocloak.GoCloak) IUserBusiness {
	return &userBusiness{
		keyCloakClient: keyCloakClient,
	}
}

func (s *userBusiness) LoginUser(ctx context.Context, credentials *requests.Credentials) (*string, error) {
	token, err := s.keyCloakClient.Login(ctx, env.Env.KeyCloakClientId, env.Env.KeyCloakClientSecret,
		env.Env.KeyCloakRealm, credentials.Username, credentials.Password)
	if err != nil {
		return nil, err
	}

	return &token.AccessToken, nil
}
