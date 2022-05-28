unit-test:
	go test -cover ./test/... -covermode=atomic

.PHONY: unit_test

lint:
	golangci-lint run --fix -v

.PHONY: lint

swag:
	swag init --parseDependency --parseInternal -g ./cmd/api/main.go -o ./docs

wire:
	wire ./internal/wired/wire.go

docker-start:
	docker-compose up --build

docker-stop:
	docker-compose down
