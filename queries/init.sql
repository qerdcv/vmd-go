CREATE TABLE IF NOT EXISTS vmd (
    id INTEGER PRIMARY KEY,
    chat_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    last_seen DATETIME,
    CONSTRAINT unq UNIQUE (chat_id, user_id)
);
CREATE TABLE IF NOT EXISTS vmd_statistics (
    id INTEGER PRIMARY KEY,
    chat_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    username VARCHAR(20) NOT NULL,
    modified_at DATETIME,
    deleted_count INTEGER DEFAULT 1,
    CONSTRAINT unq UNIQUE (chat_id, user_id)
);