-- Расписание
CREATE TABLE IF NOT EXISTS schedule (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(100) NOT NULL,
    group_id INT,
    discipline VARCHAR(255) NOT NULL,
    teacher VARCHAR(255),
    teacher_id INT REFERENCES users(id),
    start_time TIME,
    end_time TIME,
    date DATE,
    day_of_week INT,
    audience VARCHAR(50),
    lesson_type VARCHAR(50),
    week_type INT,
    comment TEXT,
    is_from_api BOOLEAN DEFAULT false,
    created_by INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Кэш расписания из API
CREATE TABLE IF NOT EXISTS schedule_cache (
    id SERIAL PRIMARY KEY,
    group_id INT,
    group_name VARCHAR(100),
    data JSONB,
    fetched_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_schedule_group ON schedule(group_name);
CREATE INDEX idx_schedule_date ON schedule(date);