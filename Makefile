# Makefile

.PHONY: build up down logs

up:
	docker-compose up -d --build

down:
	docker-compose down

lint:
	@golangci-lint run ./...

# Запуск всех тестов
test:
	go test -tags mock,integration -race -cover ./...

# Запуск всех тестов с выключенным кешированием результата
test-no-cache:
	go test -tags mock,integration -race -cover -count=1 ./...