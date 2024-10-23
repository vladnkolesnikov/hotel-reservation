build:
	@go build -o ./bin/api

run:build
	DB_NAME="hotel-reservation" \
	DB_URI="mongodb://localhost:27017" \
	go run main.go --port 4000

seed:
	go run scripts/seed.go

test:
	go test -v ./...