-- Filename: migrations/000001_create_todo_table_up.sql

CREATE TABLE IF NOT EXISTS items (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    description text NOT NULL,
    status text NOT NULL,
    mode text NOT NULL
);