#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Очистка проекта${NC}"
echo -e "${GREEN}========================================${NC}"

read -p "Вы уверены, что хотите очистить проект? (y/n): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${YELLOW}Операция отменена${NC}"
    exit 0
fi

echo -e "${YELLOW}Удаление временных файлов...${NC}"
rm -rf tmp/
rm -rf logs/*.log
rm -rf uploads/*
rm -f coverage.out
rm -f *.exe

echo -e "${YELLOW}Очистка кэша Go...${NC}"
go clean -cache
go clean -testcache

echo -e "${GREEN}✓ Проект очищен${NC}"