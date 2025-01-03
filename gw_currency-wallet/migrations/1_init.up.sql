CREATE TABLE IF NOT EXISTS users (
    id        SERIAL       PRIMARY KEY,
    username  VARCHAR(255) NOT NULL UNIQUE,
    pass_hash VARCHAR(255) NOT NULL,
    email     VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS user_balances (
    user_id INT            PRIMARY KEY,
    usd     NUMERIC(15, 2) NOT NULL DEFAULT 0.00,
    rub     NUMERIC(15, 2) NOT NULL DEFAULT 0.00,
    eur     NUMERIC(15, 2) NOT NULL DEFAULT 0.00,
    CONSTRAINT fk_user     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);