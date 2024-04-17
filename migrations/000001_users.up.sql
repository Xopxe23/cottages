CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    password VARCHAR(255),
    is_active INTEGER DEFAULT 1,
    is_verified INTEGER DEFAULT 0,
    is_staff INTEGER DEFAULT 0
);