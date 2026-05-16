#!/bin/bash

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Заполнение базы данных тестовыми данными${NC}"
echo -e "${GREEN}========================================${NC}"

# Загрузка переменных окружения
if [ -f .env ]; then
    source .env
fi

# Параметры подключения к БД
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}
DB_NAME=${DB_NAME:-student_app}

# Проверка подключения
echo -e "${YELLOW}Проверка подключения...${NC}"
if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "SELECT 1" &> /dev/null; then
    echo -e "${RED}✗ Не удалось подключиться к базе данных${NC}"
    exit 1
fi

# Вставка тестовых пользователей
echo -e "${YELLOW}Добавление тестовых пользователей...${NC}"

PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME << EOF
-- Администратор (пароль: admin123 захеширован)
INSERT INTO users (login, password, name, role) VALUES 
('admin', '\$2a\$10\$N9qo8uLOickgx2ZMRZoMy.Mr/FqMq9jY7jQqH8LrQX3yJjXxQxQxK', 'Администратор', 'admin')
ON CONFLICT (login) DO NOTHING;

-- Преподаватель (пароль: teacher123)
INSERT INTO users (login, password, name, role, department, position) VALUES 
('teacher', '\$2a\$10\$N9qo8uLOickgx2ZMRZoMy.Mr/FqMq9jY7jQqH8LrQX3yJjXxQxQxK', 'Иванов Иван Иванович', 'teacher', 'ИВИТШ', 'Старший преподаватель')
ON CONFLICT (login) DO NOTHING;

-- Сотрудник (пароль: staff123)
INSERT INTO users (login, password, name, role) VALUES 
('staff', '\$2a\$10\$N9qo8uLOickgx2ZMRZoMy.Mr/FqMq9jY7jQqH8LrQX3yJjXxQxQxK', 'Петров Петр Петрович', 'staff')
ON CONFLICT (login) DO NOTHING;

-- Студент (пароль: student123)
INSERT INTO users (login, password, name, role, group_name, course, faculty) VALUES 
('23-ПМбо-014', '\$2a\$10\$N9qo8uLOickgx2ZMRZoMy.Mr/FqMq9jY7jQqH8LrQX3yJjXxQxQxK', 'Иванов Иван Иванович', 'student', '23-ПМбо-1', 3, 'ИВИТШ')
ON CONFLICT (login) DO NOTHING;

-- Еще студенты
INSERT INTO users (login, password, name, role, group_name, course, faculty) VALUES 
('23-ПМбо-001', '\$2a\$10\$N9qo8uLOickgx2ZMRZoMy.Mr/FqMq9jY7jQqH8LrQX3yJjXxQxQxK', 'Сидоров Сидор Сидорович', 'student', '23-ПМбо-1', 3, 'ИВИТШ'),
('23-ПМбо-002', '\$2a\$10\$N9qo8uLOickgx2ZMRZoMy.Mr/FqMq9jY7jQqH8LrQX3yJjXxQxQxK', 'Петрова Анна Ивановна', 'student', '23-ПМбо-1', 3, 'ИВИТШ'),
('23-ИСбо-001', '\$2a\$10\$N9qo8uLOickgx2ZMRZoMy.Mr/FqMq9jY7jQqH8LrQX3yJjXxQxQxK', 'Козлов Дмитрий Алексеевич', 'student', '23-ИСбо-1', 3, 'ИВИТШ')
ON CONFLICT (login) DO NOTHING;

-- Профили студентов
INSERT INTO student_profiles (user_id, hobbies, clubs) 
SELECT id, ARRAY['Программирование', 'Спорт', 'Чтение'], ARRAY['Студенческий совет', 'Киберспорт']
FROM users WHERE login = '23-ПМбо-014'
ON CONFLICT (user_id) DO NOTHING;

-- Тестовые ачивки
INSERT INTO achievements (title, description, xp, icon, rarity, category) VALUES 
('🌙 Полуночник', 'Сдать домашнее задание после 2 часов ночи', 50, '🌙', 'common', 'Храбрецы'),
('⚡ Спринтер', 'Сдать задание за 30 минут до начала пары', 40, '⚡', 'common', 'Храбрецы'),
('🎯 Снайпер', 'Получить 5 отличных оценок подряд', 100, '🎯', 'rare', 'Отличники'),
('📚 Книжный червь', 'Посетить библиотеку 3 раза за неделю', 30, '📚', 'common', 'Активные'),
('🏆 Идеальная неделя', 'Сдать все домашние задания вовремя', 150, '🏆', 'epic', 'Ответственные')
ON CONFLICT (id) DO NOTHING;

-- Тестовые вопросы для экзамена
INSERT INTO exam_questions (question, answer, difficulty, subject) VALUES 
('Что такое переменная?', 'Именованная область памяти для хранения данных', 'easy', 'Программирование'),
('Что такое цикл for?', 'Конструкция для повторения блока кода определенное количество раз', 'easy', 'Программирование'),
('Что такое функция?', 'Блок кода, который можно вызывать многократно', 'medium', 'Программирование'),
('Что такое рекурсия?', 'Функция, вызывающая саму себя', 'hard', 'Программирование')
ON CONFLICT (id) DO NOTHING;

-- Тестовые пункты чек-листа
INSERT INTO checklist_items (title, description, category, xp, sort_order) VALUES 
('🎓 Получить студенческий билет', 'Обратиться в деканат', 'Документы', 50, 1),
('📚 Записаться в библиотеку', 'Прийти с фото и студенческим', 'Инфраструктура', 30, 2),
('💬 Найти свою группу в чате', 'Присоединиться к общему чату', 'Общение', 20, 3),
('📱 Скачать все приложения', 'Moodle, Teams и др.', 'Техника', 40, 4),
('🏛️ Найти деканат', 'Ориентироваться в корпусе', 'Ориентация', 15, 5)
ON CONFLICT (id) DO NOTHING;

-- Тестовые навыки
INSERT INTO skills (name, description, xp_cost, icon, category) VALUES 
('🐍 Python (базовый)', 'Переменные, циклы, функции', 50, '🐍', 'programming'),
('🐍 Python (продвинутый)', 'ООП, исключения, декораторы', 100, '🐍', 'programming'),
('📦 Git и GitHub', 'Ветки, коммиты, PR', 80, '📦', 'programming'),
('🗄️ SQL', 'SELECT, JOIN, GROUP BY', 75, '🗄️', 'programming')
ON CONFLICT (id) DO NOTHING;

EOF

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Тестовые данные успешно добавлены${NC}"
else
    echo -e "${RED}✗ Ошибка при добавлении данных${NC}"
    exit 1
fi

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Готово!${NC}"