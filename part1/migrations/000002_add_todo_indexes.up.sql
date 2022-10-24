-- Filename: migrations/000003_add_todo_indexes.up.sql

CREATE INDEX IF NOT EXISTS items_name_idx ON items USING GIN(to_tsvector('simple', name));
CREATE INDEX IF NOT EXISTS items_description_idx ON items USING GIN(to_tsvector('simple', description));
CREATE INDEX IF NOT EXISTS items_status_idx ON items USING GIN(to_tsvector('simple', status));
CREATE INDEX IF NOT EXISTS items_mode_idx ON items USING GIN(mode);