.PHONY: run build migrate seed test clean

# Переменные
BINARY_NAME=student-app
MIGRATE=migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/student_app?sslmode=disable"

# Запуск сервера
run:
	go run cmd/api/main.go

# Сборка бинарника
build:
	go build -o $(BINARY_NAME) cmd/api/main.go

# Запуск миграций
migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

# Заполнение тестовыми данными
seed:
	go run cmd/seed/main.go

# Запуск тестов
test:
	go test -v ./...

# Очистка
clean:
	rm -f $(BINARY_NAME)
	rm -rf logs/
	rm -rf uploads/

# Линтер
lint:
	golangci-lint run

# Форматирование кода
fmt:
	go fmt ./...

# Все проверки
check: fmt lint test

# Docker сборка
docker-build:
	docker build -t student-app -f docker/Dockerfile .

docker-up:
	docker-compose -f docker/docker-compose.yml up -d

docker-down:
	docker-compose -f docker/docker-compose.yml down

# Установка инструментов
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Помощь
help:
	@echo "Доступные команды:"
	@echo "  make run         - Запустить сервер"
	@echo "  make build       - Собрать бинарник"
	@echo "  make migrate-up  - Применить миграции"
	@echo "  make migrate-down- Откатить миграции"
	@echo "  make seed        - Заполнить тестовыми данными"
	@echo "  make test        - Запустить тесты"
	@echo "  make clean       - Очистить"
	@echo "  make fmt         - Форматировать код"
	@echo "  make docker-up   - Запустить Docker контейнеры"
	@echo "  make docker-down - Остановить Docker контейнеры"