-- Дерево навыков
CREATE TABLE IF NOT EXISTS skills (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    xp_cost INT DEFAULT 0,
    icon VARCHAR(10),
    category VARCHAR(100),
    parent_skill_id INT REFERENCES skills(id),
    required_skill_id INT REFERENCES skills(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Открытые навыки пользователей
CREATE TABLE IF NOT EXISTS user_skills (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    skill_id INT REFERENCES skills(id) ON DELETE CASCADE,
    unlocked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, skill_id)
);