#!/bin/bash

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Конфигурация
BACKUP_DIR="./backups"
DATE=$(date +"%Y%m%d_%H%M%S")
BACKUP_FILE="$BACKUP_DIR/student_app_backup_$DATE.sql"

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Резервное копирование базы данных${NC}"
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

# Создание папки для бэкапов
mkdir -p $BACKUP_DIR

# Проверка подключения
echo -e "${YELLOW}Проверка подключения...${NC}"
if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "SELECT 1" &> /dev/null; then
    echo -e "${RED}✗ Не удалось подключиться к базе данных${NC}"
    exit 1
fi

# Создание бэкапа
echo -e "${YELLOW}Создание бэкапа...${NC}"
PGPASSWORD=$DB_PASSWORD pg_dump -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -F c -f $BACKUP_FILE

if [ $? -eq 0 ]; then
    # Получаем размер файла
    SIZE=$(du -h $BACKUP_FILE | cut -f1)
    echo -e "${GREEN}✓ Бэкап создан: $BACKUP_FILE ($SIZE)${NC}"
    
    # Удаляем старые бэкапы (старше 30 дней)
    find $BACKUP_DIR -name "*.sql" -type f -mtime +30 -delete
    echo -e "${GREEN}✓ Старые бэкапы удалены${NC}"
else
    echo -e "${RED}✗ Ошибка при создании бэкапа${NC}"
    exit 1
fi

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Готово!${NC}"