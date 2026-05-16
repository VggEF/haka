#!/bin/bash

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Запуск тестов${NC}"
echo -e "${GREEN}========================================${NC}"

# Загрузка переменных окружения для тестов
export TEST_MODE=true
export DB_NAME=student_app_test

# Создание тестовой базы данных
echo -e "${YELLOW}Создание тестовой БД...${NC}"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c "DROP DATABASE IF EXISTS student_app_test"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c "CREATE DATABASE student_app_test"

# Применение миграций на тестовую БД
echo -e "${YELLOW}Применение миграций...${NC}"
DATABASE_URL="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/student_app_test?sslmode=disable"
migrate -path ./migrations -database "$DATABASE_URL" up

# Запуск тестов
echo -e "${YELLOW}Запуск unit тестов...${NC}"
go test -v -cover ./internal/... -coverprofile=coverage.out

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Все тесты пройдены${NC}"
    
    # Показать покрытие
    go tool cover -func=coverage.out | tail -n 1
else
    echo -e "${RED}✗ Некоторые тесты не пройдены${NC}"
    exit 1
fi

# Очистка
echo -e "${YELLOW}Очистка...${NC}"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c "DROP DATABASE IF EXISTS student_app_test"
rm -f coverage.out

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Готово!${NC}"