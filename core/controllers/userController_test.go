package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TulioGuaraldoB/cross-roads/core/businesses/mock"
	"github.com/TulioGuaraldoB/cross-roads/core/controllers"
	"github.com/TulioGuaraldoB/cross-roads/core/dtos/requests"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type userControllerTest struct {
	description         string
	setMocks            func(*mock.MockIUserBusiness)
	expectedGinContext  *gin.Context
	expectedGinRecorder *httptest.ResponseRecorder
	expectedBodyReader  io.Reader
	expectedStatusCode  int
}

func mockGinContext(httpMethod string, bodyReader io.Reader) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(httpMethod, "/", bodyReader)
	ctx.Request = req
	return ctx, w
}

func mockCredentials() *requests.Credentials {
	mockCredentials := new(requests.Credentials)
	mockCredentials.Username = "joe"
	mockCredentials.Password = "123456"

	return mockCredentials
}

func mockUserRequest() *requests.UserRequest {
	mockUserRequest := new(requests.UserRequest)
	mockUserRequest.FirstName = "Frodo"
	mockUserRequest.LastName = "Baggings"
	mockUserRequest.Email = "frodo.baggins@hotmail.com"
	mockUserRequest.Username = "frodo"
	mockUserRequest.Password = "sword"

	return mockUserRequest
}

func TestLoginController(t *testing.T) {
	mockCredentialsJSON, _ := json.Marshal(mockCredentials())
	mockCredentialsBuffer := bytes.NewBuffer(mockCredentialsJSON)

	mockContext, mockRecorder := mockGinContext(http.MethodPost, mockCredentialsBuffer)
	mockFailLoginContext, mockFailRecorder := mockGinContext(http.MethodPost, nil)
	mockToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	tests := []userControllerTest{
		{
			description: "Should return StatusCode 200 (OK)",
			setMocks: func(mib *mock.MockIUserBusiness) {
				mib.EXPECT().
					LoginUser(mockContext, mockCredentials()).
					Return(&mockToken, nil)
			},
			expectedGinContext:  mockContext,
			expectedGinRecorder: mockRecorder,
			expectedBodyReader:  mockCredentialsBuffer,
			expectedStatusCode:  http.StatusOK,
		},
		{
			description:         "Should return StatusCode 400 (Bad Request)",
			setMocks:            func(mib *mock.MockIUserBusiness) {},
			expectedGinContext:  mockFailLoginContext,
			expectedGinRecorder: mockFailRecorder,
			expectedBodyReader:  nil,
			expectedStatusCode:  http.StatusBadRequest,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			mub := mock.NewMockIUserBusiness(controller)
			testCase.setMocks(mub)

			ctx := testCase.expectedGinContext
			recorder := testCase.expectedGinRecorder

			// Act
			userController := controllers.NewUserController(mub)
			userController.Login(ctx)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, recorder.Result().StatusCode)
		})
	}
}

func TestLoginUnauthorized(t *testing.T) {
	mockCredentialsJSON, _ := json.Marshal(mockCredentials())
	mockCredentialsBuffer := bytes.NewBuffer(mockCredentialsJSON)

	mockContext, mockRecorder := mockGinContext(http.MethodPost, mockCredentialsBuffer)

	tests := []userControllerTest{
		{
			description: "Should return StatusCode 401 (Unauthorized)",
			setMocks: func(mib *mock.MockIUserBusiness) {
				mib.EXPECT().
					LoginUser(mockContext, mockCredentials()).
					Return(nil, errors.New("Invalid login"))
			},
			expectedGinContext:  mockContext,
			expectedGinRecorder: mockRecorder,
			expectedBodyReader:  mockCredentialsBuffer,
			expectedStatusCode:  http.StatusUnauthorized,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			mub := mock.NewMockIUserBusiness(controller)
			testCase.setMocks(mub)

			ctx := testCase.expectedGinContext
			recorder := testCase.expectedGinRecorder

			// Act
			userController := controllers.NewUserController(mub)
			userController.Login(ctx)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, recorder.Result().StatusCode)
		})
	}
}

func TestRegisterUnauthorized(t *testing.T) {
	mockUserRequestJSON, _ := json.Marshal(mockUserRequest())
	mockCredentialsBuffer := bytes.NewBuffer(mockUserRequestJSON)

	mockContext, mockRecorder := mockGinContext(http.MethodPost, mockCredentialsBuffer)

	tests := []userControllerTest{
		{
			description: "Should return StatusCode 401 (Unauthorized)",
			setMocks: func(mib *mock.MockIUserBusiness) {
				mib.EXPECT().
					RegisterUser(mockContext, mockUserRequest()).
					Return(nil, errors.New("Invalid user credentials"))
			},
			expectedGinContext:  mockContext,
			expectedGinRecorder: mockRecorder,
			expectedBodyReader:  mockCredentialsBuffer,
			expectedStatusCode:  http.StatusUnauthorized,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			mub := mock.NewMockIUserBusiness(controller)
			testCase.setMocks(mub)

			ctx := testCase.expectedGinContext
			recorder := testCase.expectedGinRecorder

			// Act
			userController := controllers.NewUserController(mub)
			userController.Register(ctx)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, recorder.Result().StatusCode)
		})
	}
}

func TestRegisterController(t *testing.T) {
	mockUserRequestJSON, _ := json.Marshal(mockUserRequest())
	mockUserRequestBuffer := bytes.NewBuffer(mockUserRequestJSON)

	mockContext, mockRecorder := mockGinContext(http.MethodPost, mockUserRequestBuffer)
	mockFailLoginContext, mockFailRecorder := mockGinContext(http.MethodPost, nil)
	mockToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	tests := []userControllerTest{
		{
			description: "Should return StatusCode 200 (OK)",
			setMocks: func(mib *mock.MockIUserBusiness) {
				mib.
					EXPECT().RegisterUser(mockContext, mockUserRequest()).
					Return(&mockToken, nil)
			},
			expectedGinContext:  mockContext,
			expectedGinRecorder: mockRecorder,
			expectedBodyReader:  mockUserRequestBuffer,
			expectedStatusCode:  http.StatusOK,
		},
		{
			description:         "Should return StatusCode 400 (Bad Request)",
			setMocks:            func(mib *mock.MockIUserBusiness) {},
			expectedGinContext:  mockFailLoginContext,
			expectedGinRecorder: mockFailRecorder,
			expectedBodyReader:  nil,
			expectedStatusCode:  http.StatusBadRequest,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			controller := gomock.NewController(t)
			defer controller.Finish()

			mub := mock.NewMockIUserBusiness(controller)
			testCase.setMocks(mub)

			ctx := testCase.expectedGinContext
			recorder := testCase.expectedGinRecorder

			// Act
			userController := controllers.NewUserController(mub)
			userController.Register(ctx)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, recorder.Result().StatusCode)
		})
	}
}
