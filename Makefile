run:
	go run ./cmd/server


swagger:
	swag init -g cmd/server/main.go -o internal/docs


tidy:
	go mod tidy