# Makefile
.PHONY: fmt lint tidy

# Format all code
fmt:
	gofmt -w .

fmt-check:
	@gofmt -l . | grep . && (echo "Some files are not formatted" && exit 1) || echo "All files are properly formatted"

lint:
	golangci-lint run

tidy:
	go mod tidy

test:
	go test ./... -v

build:
	go build -v ./...-

run: 
	go run ./cmd/main.go

create-migration:
	migrate create -ext sql -dir ./migrations -seq $(name)

swagger:
	swag init -g cmd/main.go