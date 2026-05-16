#!/bin/bash

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Запуск StudentApp в режиме разработки${NC}"
echo -e "${GREEN}========================================${NC}"

# Загрузка переменных окружения
if [ -f .env ]; then
    source .env
    echo -e "${GREEN}✓ Загружены переменные окружения${NC}"
else
    echo -e "${YELLOW}⚠ Файл .env не найден, используем значения по умолчанию${NC}"
fi

# Создание необходимых папок
mkdir -p logs
mkdir -p uploads
mkdir -p tmp

# Проверка наличия Docker (опционально)
if command -v docker &> /dev/null; then
    echo -e "${GREEN}✓ Docker найден${NC}"
else
    echo -e "${YELLOW}⚠ Docker не найден (опционально)${NC}"
fi

# Запуск миграций
echo -e "${YELLOW}Применение миграций...${NC}"
./scripts/migrate.sh

# Заполнение тестовыми данными
echo -e "${YELLOW}Заполнение тестовыми данными...${NC}"
./scripts/seed.sh

# Установка зависимостей
echo -e "${YELLOW}Установка зависимостей...${NC}"
go mod download
go mod tidy

# Запуск сервера
echo -e "${GREEN}Запуск сервера...${NC}"
echo -e "${GREEN}Сервер запущен на http://localhost:${SERVER_PORT:-8080}${NC}"
echo -e "${GREEN}Swagger документация: http://localhost:${SERVER_PORT:-8080}/swagger/index.html${NC}"
echo -e "${YELLOW}Нажмите Ctrl+C для остановки${NC}"

# Проверка наличия air для hot reload
if command -v air &> /dev/null; then
    echo -e "${GREEN}✓ Запуск с hot reload (air)${NC}"
    air
else
    echo -e "${YELLOW}⚠ air не найден, устанавливаем...${NC}"
    go install github.com/cosmtrek/air@latest
    air
fi