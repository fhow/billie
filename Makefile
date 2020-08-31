include ./build/.env
export $(shell sed 's/=.*//' ./build/.env)

deps: tidy vendor

## Dependencies
vendor:
	go mod vendor

tidy:
	go mod tidy

run:
	@echo "Running app..."
	@go run cmd/main.go

tests:
	@echo "Running functional tests..."
	@go test -mod=vendor ./internal/tests/tests/... -cover -count=1 -short

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run

init: deps run