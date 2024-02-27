-- +goose Up
-- +goose StatementBegin
INSERT INTO roles (name)
VALUES ('Admin'),
       ('Moderator'),
       ('User');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM roles
WHERE name IN ('Admin', 'Moderator', 'User');
-- +goose StatementEnd
