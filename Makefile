# Makefile
.PHONY: fmt lint tidy

# Format all code
fmt:
	gofmt -w .

fmt-check:
	@echo "Checking formatting..."
	@gofmt -l . | grep . && (echo "Some files are not formatted" && exit 1) || echo "All files are properly formatted"

lint:
	@echo "Running linter..."
	golangci-lint run

tidy:
	go mod tidy

test:
	@echo "Running tests..."
	go test ./... -v

build:
	@echo "Building the project..."
	go build -v ./...-

run: 
	go run ./cmd/main.go

create-migration:
	migrate create -ext sql -dir ./migrations -seq $(name)

swagger:
	swag init -g cmd/main.go