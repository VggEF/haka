-- ДПО и мероприятия
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    short_text TEXT,
    full_text TEXT,
    date DATE,
    time TIME,
    type VARCHAR(50),
    category VARCHAR(100),
    location VARCHAR(255),
    price VARCHAR(100),
    organizer VARCHAR(255),
    image_url TEXT,
    available_spots INT,
    contact VARCHAR(255),
    registrations JSONB DEFAULT '[]',
    created_by INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Регистрации на мероприятия
CREATE TABLE IF NOT EXISTS event_registrations (
    id SERIAL PRIMARY KEY,
    event_id INT REFERENCES events(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50) DEFAULT 'registered',
    UNIQUE(event_id, user_id)
);

CREATE INDEX idx_events_date ON events(date);