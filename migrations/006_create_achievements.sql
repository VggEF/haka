-- Ачивки
CREATE TABLE IF NOT EXISTS achievements (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    xp INT DEFAULT 0,
    icon VARCHAR(10),
    rarity VARCHAR(50),
    category VARCHAR(100),
    condition_text TEXT,
    created_by INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Полученные ачивки
CREATE TABLE IF NOT EXISTS user_achievements (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    achievement_id INT REFERENCES achievements(id) ON DELETE CASCADE,
    earned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, achievement_id)
);

-- Ежедневные задания
CREATE TABLE IF NOT EXISTS daily_quests (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    xp INT DEFAULT 0,
    coins INT DEFAULT 0,
    requirement TEXT,
    expires_at DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Выполненные задания
CREATE TABLE IF NOT EXISTS user_quests (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    quest_id INT REFERENCES daily_quests(id) ON DELETE CASCADE,
    completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, quest_id)
);