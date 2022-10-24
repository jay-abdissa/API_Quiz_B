-- Filename: migrations/000002_add_todo_indexes.down.sql
DROP INDEX IF EXISTS items_name_idx;
DROP INDEX IF EXISTS items_description_idx;
DROP INDEX IF EXISTS items_status_idx;
DROP INDEX IF EXISTS items_mode_idx;