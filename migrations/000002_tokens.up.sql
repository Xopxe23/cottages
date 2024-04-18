CREATE TABLE refresh_tokens (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(64) NOT NULL,
    expires_at TIMESTAMP NOT NULL
);