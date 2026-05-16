-- Чек-лист
CREATE TABLE IF NOT EXISTS checklist_items (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(100),
    xp INT DEFAULT 0,
    sort_order INT DEFAULT 0,
    is_active BOOLEAN DEFAULT true
);

-- Прогресс чек-листа
CREATE TABLE IF NOT EXISTS user_checklist (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    item_id INT REFERENCES checklist_items(id) ON DELETE CASCADE,
    completed_at TIMESTAMP,
    is_completed BOOLEAN DEFAULT false,
    PRIMARY KEY (user_id, item_id)
);