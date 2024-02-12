-- +goose Up
-- +goose StatementBegin
CREATE TABLE test (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE test;
-- +goose StatementEnd
