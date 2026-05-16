#!/bin/bash

GREEN='\033[0;32m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Остановка Docker контейнеров${NC}"
echo -e "${GREEN}========================================${NC}"

if ! command -v docker-compose &> /dev/null; then
    docker compose -f docker/docker-compose.yml down
else
    docker-compose -f docker/docker-compose.yml down
fi

echo -e "${GREEN}✓ Контейнеры остановлены${NC}"