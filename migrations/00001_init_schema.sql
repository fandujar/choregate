-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    slug TEXT NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    system_role TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    namespace TEXT NOT NULL,
    description TEXT,
    timeout INTERVAL,
    task_scope JSONB,
    task_spec JSONB
);

CREATE TABLE IF NOT EXISTS task_runs (
    id UUID PRIMARY KEY,
    task_id UUID NOT NULL,
    status JSONB NOT NULL,
    task_run JSONB NOT NULL
);

-- +goose Down
-- SQL in section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS task_runs;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS users;
