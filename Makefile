all: build run

build:
	@go build -o "./" ./cmd/api/main.go

run:
	./main