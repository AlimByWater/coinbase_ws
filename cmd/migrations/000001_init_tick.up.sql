CREATE TABLE IF NOT EXISTS tick (
    id SERIAL PRIMARY KEY,
    symbol TEXT,
    bid FLOAT(50),
    ask FLOAT(50),
    created_at TEXT
);