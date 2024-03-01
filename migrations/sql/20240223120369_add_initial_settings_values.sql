-- +goose Up
-- +goose StatementBegin
INSERT INTO settings (name, value)
VALUES ('default_role_id', '3');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM settings
WHERE name = 'default_role_id';
-- +goose StatementEnd
