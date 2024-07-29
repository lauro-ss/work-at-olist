all: build run

build:
	@go build -o "./" ./cmd/api/main.go

run:
	./main

swag:
	@swag init -d ./cmd/api/ --pdl 3