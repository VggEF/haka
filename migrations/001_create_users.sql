-- Пользователи
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'student',
    group_name VARCHAR(100),
    course INT DEFAULT 1,
    faculty VARCHAR(100),
    photo TEXT,
    email VARCHAR(255),
    phone VARCHAR(50),
    telegram VARCHAR(100),
    vk VARCHAR(255),
    github VARCHAR(255),
    total_xp INT DEFAULT 0,
    coins INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Профили студентов
CREATE TABLE IF NOT EXISTS student_profiles (
    user_id INT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    hobbies TEXT[],
    clubs TEXT[],
    about TEXT,
    checklist_items JSONB DEFAULT '[]',
    planner_data JSONB DEFAULT '{}'
);

-- Профили преподавателей
CREATE TABLE IF NOT EXISTS teacher_profiles (
    user_id INT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    department VARCHAR(255),
    position VARCHAR(255),
    degree VARCHAR(255),
    office VARCHAR(50),
    experience INT DEFAULT 0,
    achievements JSONB DEFAULT '[]'
);

-- Индексы
CREATE INDEX idx_users_login ON users(login);
CREATE INDEX idx_users_role ON users(role);