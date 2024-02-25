-- +goose Up
-- +goose StatementBegin
CREATE TABLE denied_tokens
(
    id         SERIAL PRIMARY KEY,
    token      VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE denied_tokens;
-- +goose StatementEnd
