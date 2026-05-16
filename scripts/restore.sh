#!/bin/bash

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Восстановление базы данных из бэкапа${NC}"
echo -e "${GREEN}========================================${NC}"

# Загрузка переменных окружения
if [ -f .env ]; then
    source .env
fi

# Параметры подключения
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}
DB_NAME=${DB_NAME:-student_app}
BACKUP_DIR="./backups"

# Список доступных бэкапов
echo -e "${YELLOW}Доступные бэкапы:${NC}"
ls -la $BACKUP_DIR/*.sql 2>/dev/null || echo "Нет доступных бэкапов"

if [ ! -d "$BACKUP_DIR" ] || [ -z "$(ls -A $BACKUP_DIR)" ]; then
    echo -e "${RED}✗ Нет доступных бэкапов для восстановления${NC}"
    exit 1
fi

# Выбор бэкапа
echo ""
read -p "Введите имя файла бэкапа для восстановления: " BACKUP_FILE

if [ ! -f "$BACKUP_DIR/$BACKUP_FILE" ]; then
    echo -e "${RED}✗ Файл не найден${NC}"
    exit 1
fi

# Предупреждение
echo -e "${RED}⚠ ВНИМАНИЕ! Восстановление удалит текущие данные!${NC}"
read -p "Вы уверены? (y/n): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${YELLOW}Операция отменена${NC}"
    exit 0
fi

# Восстановление
echo -e "${YELLOW}Восстановление из бэкапа $BACKUP_FILE...${NC}"
PGPASSWORD=$DB_PASSWORD pg_restore -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c -v $BACKUP_DIR/$BACKUP_FILE

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ База данных успешно восстановлена${NC}"
else
    echo -e "${RED}✗ Ошибка при восстановлении${NC}"
    exit 1
fi

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Готово!${NC}"