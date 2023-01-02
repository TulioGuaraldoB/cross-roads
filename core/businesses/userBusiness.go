package businesses

import (
	"context"

	"github.com/TulioGuaraldoB/cross-roads/core/dtos/requests"
	"github.com/TulioGuaraldoB/cross-roads/infra/services"
)

type IUserBusiness interface {
	LoginUser(ctx context.Context, credentials *requests.Credentials) (*string, error)
}

type userBusiness struct {
	keyCloakService services.IKeyCloakService
}

func NewUserBusinesses(keyCloakService services.IKeyCloakService) IUserBusiness {
	return &userBusiness{
		keyCloakService: keyCloakService,
	}
}

func (s *userBusiness) LoginUser(ctx context.Context, credentials *requests.Credentials) (*string, error) {
	return s.keyCloakService.Login(ctx, credentials)
}
