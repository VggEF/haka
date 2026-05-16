-- Библиотека
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255),
    publisher VARCHAR(255),
    year VARCHAR(10),
    isbn VARCHAR(50),
    description TEXT,
    cover_url TEXT,
    file_path VARCHAR(500),
    category VARCHAR(100),
    tags TEXT[],
    available_copies INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Кэш внешней библиотеки
CREATE TABLE IF NOT EXISTS library_cache (
    id SERIAL PRIMARY KEY,
    search_query VARCHAR(255),
    data JSONB,
    fetched_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);