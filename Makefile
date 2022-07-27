
ask oppo:
	go run main.go ask --type=opposite

ask syno:
	go run main.go ask --type=synonym

ask tran:
	go run main.go ask --type=translation

lint:
	golangci-lint run

unit-test:
	go clean -testcache && go test ./... -short

mockgen:
	mockgen -destination=internal/mocks/login/mock_login_client.go -package mocks github.com/eneskzlcn/dictionary-app-cli/internal/login LoginClient
	mockgen -destination=internal/mocks/login/mock_login_service.go -package mocks github.com/eneskzlcn/dictionary-app-cli/internal/login LoginService