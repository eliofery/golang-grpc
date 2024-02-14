-- +goose Up
-- +goose StatementBegin
CREATE TABLE test (
    id serial primary key,
    name varchar(255),
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE test;
-- +goose StatementEnd
