-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         serial primary key,
    first_name varchar(50),
    last_name  varchar(50),
    email      varchar(255) not null unique,
    password   varchar(255) not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
