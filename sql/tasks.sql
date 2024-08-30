CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    namespace TEXT NOT NULL,
    description TEXT,
    timeout INTERVAL,
    task_scope JSONB,
    task_spec JSONB
);