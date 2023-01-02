package businesses_test

import (
	"context"
	"testing"

	"github.com/TulioGuaraldoB/cross-roads/core/businesses"
	"github.com/TulioGuaraldoB/cross-roads/core/dtos/requests"
	"github.com/TulioGuaraldoB/cross-roads/infra/services/mock"
	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type userBusinessTest struct {
	description         string
	setMocks            func(*mock.MockIKeyCloakService)
	expectedContext     context.Context
	expectedCredentials *requests.Credentials
}

func TestUserBusinessLogin(t *testing.T) {
	mockContext := context.Background()
	mockCredentials := new(requests.Credentials)
	faker.Struct(mockCredentials)

	mockToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	tests := []userBusinessTest{
		{
			description: "Should return no error on user login",
			setMocks: func(mikcs *mock.MockIKeyCloakService) {
				mikcs.EXPECT().
					Login(mockContext, mockCredentials).
					Return(&mockToken, nil)
			},
			expectedContext:     mockContext,
			expectedCredentials: mockCredentials,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			mks := mock.NewMockIKeyCloakService(controller)
			testCase.setMocks(mks)

			// Act
			userBusiness := businesses.NewUserBusinesses(mks)
			token, err := userBusiness.LoginUser(testCase.expectedContext, testCase.expectedCredentials)
			if err != nil { // Assert
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
			assert.NotEmpty(t, token)
		})
	}
}
