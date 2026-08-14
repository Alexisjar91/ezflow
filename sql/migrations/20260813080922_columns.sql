-- +goose Up
CREATE TABLE IF NOT EXISTS columns (
    id CHAR(26) PRIMARY KEY,
    project_id CHAR(26) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    color VARCHAR(10) NOT NULL DEFAULT '#FF6B6B',
    position INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS columns;
