-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name  VARCHAR(50),
    email      VARCHAR(128) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    role_id    INTEGER,
    created_at TIMESTAMP    NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles (id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
