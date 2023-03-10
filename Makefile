run:
	@echo Running...
	@go run main.go

install:
	@echo Downloading dependencies...
	@go get
	@echo Validating dependencies...
	@go mod tidy

vendor:
	@echo Generating vendor from dependencies...
	@go mod vendor

mock:
	@echo Generating mocks...
	@echo Mocking services...
	@mockgen -source=infra/services/keycloakService.go -destination=infra/services/mock/mockKeycloakService.go -package=mock
	@echo Mocking businesses...
	@mockgen -source=core/businesses/userBusiness.go -destination=core/businesses/mock/mockUserBusiness.go -package=mock
	@echo Mocking controllers...
	@mockgen -source=core/controllers/userController.go -destination=core/controllers/mock/mockUserController.go -package=mock

test:
	@echo Running tests...
	@go test -v ./...

coverage:
	@echo Running test coverage...
	@go test -v ./... -coverprofile=cover.out
	@go tool cover -html=cover.out -o cover.html
	@echo Coverage test successfully completed.