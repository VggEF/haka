-- Новости
CREATE TABLE IF NOT EXISTS news (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    short_text TEXT,
    full_text TEXT,
    image_url TEXT,
    category VARCHAR(100),
    date DATE DEFAULT CURRENT_DATE,
    is_pinned BOOLEAN DEFAULT false,
    created_by INT REFERENCES users(id),
    views INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_news_date ON news(date);