
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