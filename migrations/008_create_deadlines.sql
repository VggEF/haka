-- Дедлайны
CREATE TABLE IF NOT EXISTS deadlines (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    subject VARCHAR(255),
    group_name VARCHAR(100),
    due_date DATE NOT NULL,
    due_time TIME,
    priority VARCHAR(50),
    description TEXT,
    status VARCHAR(50) DEFAULT 'pending',
    created_by INT REFERENCES users(id),
    assigned_to INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_deadlines_due_date ON deadlines(due_date);