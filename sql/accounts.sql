CREATE TABLE IF NOT EXISTS accounts (
    id                uuid        NOT NULL PRIMARY KEY,
    pfp_file          text        NOT NULL,
    name              text        NOT NULL,
    login             text        NOT NULL UNIQUE,
    email             text        NOT NULL UNIQUE,
    email_verified    bool        NOT NULL DEFAULT false,
    password_hash     text        NOT NULL,
    registration_date timestamptz NOT NULL DEFAULT now(),
    two_factor        bool        NOT NULL DEFAULT false,
    role              text        NOT NULL DEFAULT 'user'
        CHECK (role IN ('user', 'admin')),
    status            text        NOT NULL DEFAULT 'active'
        CHECK (status IN ('active', 'banned', 'deleted')),
    plan_id           INT         NOT NULL REFERENCES plans(id),
    public_key        text        NOT NULL
);

CREATE INDEX idx_accounts_plan_id ON accounts(plan_id);