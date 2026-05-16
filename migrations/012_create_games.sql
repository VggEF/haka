-- Мини-игры
CREATE TABLE IF NOT EXISTS mini_games (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    icon VARCHAR(10),
    xp_reward INT DEFAULT 10,
    coin_reward INT DEFAULT 5,
    game_data JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Результаты игр
CREATE TABLE IF NOT EXISTS game_results (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    game_id INT REFERENCES mini_games(id) ON DELETE CASCADE,
    score INT,
    xp_earned INT,
    coins_earned INT,
    played_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);