CREATE TABLE IF NOT EXISTS task_runs (
    id UUID PRIMARY KEY,
    task_id UUID NOT NULL,
    status JSONB NOT NULL,
    spec JSONB NOT NULL
);