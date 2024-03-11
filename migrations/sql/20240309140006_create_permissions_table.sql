-- +goose Up
-- +goose StatementBegin
CREATE TABLE permissions
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE permissions;
-- +goose StatementEnd
