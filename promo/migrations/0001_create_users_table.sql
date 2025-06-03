-- +goose Up
-- +goose StatementBegin
CREATE TABLE clients (
    id TEXT PRIMARY KEY,
    registered_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE views (
    id TEXT PRIMARY KEY,
    client_id TEXT NOT NULL,
    entity_id TEXT NOT NULL,
    viewed_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE clicks (
    id TEXT PRIMARY KEY,
    client_id TEXT NOT NULL,
    entity_id TEXT NOT NULL,
    clicked_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE comments (
    comment_id TEXT PRIMARY KEY,
    client_id TEXT NOT NULL,
    entity_id TEXT NOT NULL,
    message TEXT NOT NULL,
    commented_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clients;
DROP TABLE views;
DROP TABLE clicks;
DROP TABLE comments;
-- +goose StatementEnd

