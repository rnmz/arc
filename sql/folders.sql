CREATE TABLE IF NOT EXISTS folders (
    id          UUID NOT NULL PRIMARY KEY,
    name        TEXT NOT NULL,
    path        TEXT NOT NULL,
    owner_id    UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    view_status INT  NOT NULL DEFAULT 32
);

CREATE INDEX idx_folders_owner_id ON folders(owner_id);