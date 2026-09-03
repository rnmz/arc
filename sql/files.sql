CREATE TABLE IF NOT EXISTS files (
    id          UUID    NOT NULL PRIMARY KEY,
    name        TEXT    NOT NULL,
    extension   TEXT    NOT NULL,
    size        BIGINT  NOT NULL CHECK (size >= 0),
    path        TEXT    NOT NULL,
    encrypt     BOOL    NOT NULL DEFAULT false,
    owner_id    UUID    NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    view_status INT     NOT NULL DEFAULT 32,
    folder_id   UUID    NOT NULL REFERENCES folders(id) ON DELETE CASCADE
);

CREATE INDEX idx_files_owner_id  ON files(owner_id);
CREATE INDEX idx_files_folder_id ON files(folder_id);