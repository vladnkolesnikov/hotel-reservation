build:
	@go build -o ./bin/api

run:build
	@go run main.go --port 4000

test:
	go test -v ./...