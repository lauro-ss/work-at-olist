all: build run

build:
	@go build -o "./" ./cmd/api/main.go

run:
	@./main

csv: build_csv run_csv

build_csv:
	@go build -o "./import_csv" ./cmd/import_csv/main.go

run_csv:
	@./import_csv authors.csv

swag:
	@swag init -d ./cmd/api/ --pdl 3

test:
	@docker compose -f docker-compose.test.yml up -d -V
	@go test ./...
	@docker compose -f docker-compose.test.yml down -v