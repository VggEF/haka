#!/bin/bash

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Сборка Docker образов${NC}"
echo -e "${GREEN}========================================${NC}"

# Сборка основного образа
echo -e "${YELLOW}Сборка образа API...${NC}"
docker build -f docker/Dockerfile -t student-app:latest .

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Образ API собран${NC}"
else
    echo -e "${RED}✗ Ошибка сборки API${NC}"
    exit 1
fi

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Готово!${NC}"