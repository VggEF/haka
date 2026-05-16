-- Привычки
CREATE TABLE IF NOT EXISTS habits (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    xp_reward INT DEFAULT 10,
    created_by INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Привычки пользователей
CREATE TABLE IF NOT EXISTS user_habits (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    habit_id INT REFERENCES habits(id) ON DELETE CASCADE,
    streak INT DEFAULT 0,
    last_completed DATE,
    PRIMARY KEY (user_id, habit_id)
);

-- Логи выполнения
CREATE TABLE IF NOT EXISTS habit_logs (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    habit_id INT REFERENCES habits(id) ON DELETE CASCADE,
    completed_date DATE DEFAULT CURRENT_DATE,
    xp_earned INT,
    PRIMARY KEY (user_id, habit_id, completed_date)
);