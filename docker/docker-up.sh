#!/bin/bash

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Запуск Docker Compose${NC}"
echo -e "${GREEN}========================================${NC}"

# Копирование env файла
if [ ! -f .env ]; then
    cp .env.docker .env
    echo -e "${YELLOW}Создан .env файл из .env.docker${NC}"
fi

# Выбор окружения
if [ "$1" == "dev" ]; then
    echo -e "${YELLOW}Запуск в режиме разработки...${NC}"
    docker-compose -f docker/docker-compose.dev.yml up -d
else
    echo -e "${YELLOW}Запуск в production режиме...${NC}"
    docker-compose -f docker/docker-compose.yml up -d
fi

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Контейнеры запущены${NC}"
    echo -e "${GREEN}API: http://localhost:8080${NC}"
    echo -e "${GREEN}Adminer: http://localhost:8081${NC}"
    echo -e "${GREEN}PgAdmin: http://localhost:5050${NC} (dev режим)"
    echo -e "${GREEN}MailHog: http://localhost:8025${NC} (dev режим)"
else
    echo -e "${RED}✗ Ошибка запуска контейнеров${NC}"
    exit 1
fi