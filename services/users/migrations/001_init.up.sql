-- users_auth: данные для логина
CREATE TABLE IF NOT EXISTS users_auth (
                                          id SERIAL PRIMARY KEY,
                                          email TEXT NOT NULL UNIQUE,
                                          password TEXT NOT NULL,
                                          created_at TIMESTAMP NOT NULL DEFAULT now()
    );

-- users: профиль пользователя
CREATE TABLE IF NOT EXISTS users (
                                     id INTEGER PRIMARY KEY,
                                     username TEXT NOT NULL,
                                     bio TEXT,
                                     avatar_url TEXT,
                                     created_at TIMESTAMP NOT NULL DEFAULT now()
    );

-- (опционально, но полезно)
CREATE INDEX IF NOT EXISTS idx_users_auth_email ON users_auth(email);