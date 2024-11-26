CREATE TABLE urls(
    id BIGINT PRIMARY KEY,
    slug VARCHAR(11) UNIQUE,
    url VARCHAR(255) UNIQUE,
    clicked INT NOT NULL DEFAULT 0,

    created_at TIMESTAMP DEFAULT now(),
    expires_at TIMESTAMP
);
