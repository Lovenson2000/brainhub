-- Users table
CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        firstname VARCHAR(100),
        lastname VARCHAR(100),
        email VARCHAR(100) UNIQUE,
        password VARCHAR(100),
        school VARCHAR(100),
        major VARCHAR(100),
        bio TEXT
    );

-- Posts table
CREATE TABLE
    posts (
        id SERIAL PRIMARY KEY,
        user_id INTEGER REFERENCES users (id),
        content TEXT,
        image VARCHAR(255),
        created_at TIMESTAMP
    );

-- Study sessions table
CREATE TABLE
    study_sessions (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255),
        description TEXT,
        start_time TIMESTAMP,
        end_time TIMESTAMP,
        location VARCHAR(255)
    );

-- Documents table (related to study sessions)
CREATE TABLE
    documents (
        id SERIAL PRIMARY KEY,
        study_session_id INTEGER REFERENCES study_sessions (id),
        title VARCHAR(255),
        url VARCHAR(255),
        uploaded_by INTEGER REFERENCES users (id),
        uploaded_at TIMESTAMP
    );

-- Chat messages table (related to study sessions)
CREATE TABLE
    chat_messages (
        id SERIAL PRIMARY KEY,
        study_session_id INTEGER REFERENCES study_sessions (id),
        user_id INTEGER REFERENCES users (id),
        message TEXT,
        created_at TIMESTAMP
    );

CREATE TABLE
    IF NOT EXISTS documents (
        id SERIAL PRIMARY KEY,
        study_session_id INTEGER REFERENCES study_sessions (id) ON DELETE CASCADE,
        user_id INTEGER REFERENCES users (id),
        url TEXT NOT NULL,
        uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    IF NOT EXISTS notes (
        id SERIAL PRIMARY KEY,
        study_session_id INTEGER REFERENCES study_sessions (id) ON DELETE CASCADE,
        user_id INTEGER REFERENCES users (id),
        content TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );