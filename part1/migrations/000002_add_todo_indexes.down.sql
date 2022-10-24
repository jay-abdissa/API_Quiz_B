-- Filename: migrations/000002_add_todo_indexes.down.sql

DROP INDEX If EXISTS items_name_idx;
DROP INDEX If EXISTS items_description_idx;
DROP INDEX If EXISTS items_status_idx;
DROP INDEX If EXISTS items_mode_idx;