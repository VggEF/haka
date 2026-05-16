#!/bin/bash

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Остановка Docker Compose${NC}"
echo -e "${GREEN}========================================${NC}"

if [ "$1" == "dev" ]; then
    docker-compose -f docker/docker-compose.dev.yml down
else
    docker-compose -f docker/docker-compose.yml down
fi

echo -e "${GREEN}✓ Контейнеры остановлены${NC}"