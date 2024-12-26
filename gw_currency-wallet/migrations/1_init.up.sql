CREATE TABLE IF NOT EXISTS users (
    id        SERIAL       PRIMARY KEY,
    username  VARCHAR(255) NOT NULL,
    pass_hash VARCHAR(255) NOT NULL,
    email     VARCHAR(255) NOT NULL UNIQUE
);