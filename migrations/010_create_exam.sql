-- Вопросы для экзамена
CREATE TABLE IF NOT EXISTS exam_questions (
    id SERIAL PRIMARY KEY,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    difficulty VARCHAR(50),
    subject VARCHAR(255),
    category VARCHAR(100),
    created_by INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Результаты экзаменов
CREATE TABLE IF NOT EXISTS exam_results (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    subject VARCHAR(255),
    score INT,
    total_questions INT,
    correct_answers INT,
    taken_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);