build:
	@go build -o ./bin/api

run:build
	DB_NAME="hotel-reservation" go run main.go --port 4000

test:
	go test -v ./...