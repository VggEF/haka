#!/bin/bash

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Запуск миграций базы данных${NC}"
echo -e "${GREEN}========================================${NC}"

# Загрузка переменных окружения
if [ -f .env ]; then
    source .env
fi

# Параметры подключения к БД
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}
DB_NAME=${DB_NAME:-student_app}

# Проверка наличия migrate
if ! command -v migrate &> /dev/null; then
    echo -e "${YELLOW}migrate не найден. Установка...${NC}"
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
fi

# Проверка подключения к БД
echo -e "${YELLOW}Проверка подключения к PostgreSQL...${NC}"
if PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c "SELECT 1" &> /dev/null; then
    echo -e "${GREEN}✓ Подключение к PostgreSQL успешно${NC}"
else
    echo -e "${RED}✗ Не удалось подключиться к PostgreSQL${NC}"
    exit 1
fi

# Создание базы данных если не существует
echo -e "${YELLOW}Проверка наличия базы данных $DB_NAME...${NC}"
if PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -lqt | cut -d \| -f 1 | grep -qw $DB_NAME; then
    echo -e "${GREEN}✓ База данных $DB_NAME уже существует${NC}"
else
    echo -e "${YELLOW}Создание базы данных $DB_NAME...${NC}"
    PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c "CREATE DATABASE $DB_NAME"
    echo -e "${GREEN}✓ База данных $DB_NAME создана${NC}"
fi

# Запуск миграций
echo -e "${YELLOW}Применение миграций...${NC}"
DATABASE_URL="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"

migrate -path ./migrations -database "$DATABASE_URL" up

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Миграции успешно применены${NC}"
else
    echo -e "${RED}✗ Ошибка при применении миграций${NC}"
    exit 1
fi

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Готово!${NC}"