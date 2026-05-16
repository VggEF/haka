#!/bin/bash

# Цвета для вывода
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Запуск Docker контейнеров${NC}"
echo -e "${GREEN}========================================${NC}"

# Проверка наличия docker-compose
if ! command -v docker-compose &> /dev/null; then
    echo -e "${YELLOW}docker-compose не найден, пробуем docker compose...${NC}"
    docker compose -f docker/docker-compose.yml up -d
else
    docker-compose -f docker/docker-compose.yml up -d
fi

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Контейнеры запущены${NC}"
    echo -e "${GREEN}PostgreSQL: localhost:5432${NC}"
    echo -e "${GREEN}Redis: localhost:6379${NC}"
else
    echo -e "${RED}✗ Ошибка при запуске контейнеров${NC}"
    exit 1
fi