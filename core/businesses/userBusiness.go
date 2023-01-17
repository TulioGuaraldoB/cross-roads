package businesses

import (
	"context"

	"github.com/TulioGuaraldoB/cross-roads/core/dtos/requests"
	"github.com/TulioGuaraldoB/cross-roads/infra/services"
)

type IUserBusiness interface {
	LoginUser(ctx context.Context, credentials *requests.Credentials) (*string, error)
	RegisterUser(ctx context.Context, user *requests.UserRequest) (*string, error)
}

type userBusiness struct {
	keyCloakService services.IKeyCloakService
}

func NewUserBusinesses(keyCloakService services.IKeyCloakService) IUserBusiness {
	return &userBusiness{
		keyCloakService: keyCloakService,
	}
}

func (b *userBusiness) LoginUser(ctx context.Context, credentials *requests.Credentials) (*string, error) {
	return b.keyCloakService.Login(ctx, credentials)
}

func (b *userBusiness) RegisterUser(ctx context.Context, user *requests.UserRequest) (*string, error) {
	return b.keyCloakService.CreateUser(ctx, user)
}
