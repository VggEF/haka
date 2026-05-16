#!/bin/bash

GREEN='\033[0;32m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Просмотр логов${NC}"
echo -e "${GREEN}========================================${NC}"

docker-compose logs -f --tail=100