-- +goose Up
-- +goose StatementBegin
INSERT INTO permissions (name, description)
VALUES ('create_settings', 'Create settings'),
       ('read_settings', 'View settings'),
       ('update_settings', 'Edit settings'),
       ('delete_settings', 'Delete settings'),

       ('create_denied_tokens', 'Create denied tokens'),
       ('read_denied_tokens', 'View denied tokens'),
       ('update_denied_tokens', 'Edit denied tokens'),
       ('delete_denied_tokens', 'Delete denied tokens'),

       ('create_roles', 'Create roles'),
       ('read_roles', 'View roles'),
       ('update_roles', 'Edit roles'),
       ('delete_roles', 'Delete roles'),

       ('create_permissions', 'Create permissions'),
       ('read_permissions', 'View permissions'),
       ('update_permissions', 'Edit permissions'),
       ('delete_permissions', 'Delete permissions'),

       ('create_users', 'Create users'),
       ('read_users', 'View users'),
       ('update_users', 'Edit users'),
       ('delete_users', 'Delete users');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM permissions
WHERE name IN (
               'create_settings', 'read_settings', 'update_settings', 'delete_settings',
               'create_denied_tokens', 'read_denied_tokens', 'update_denied_tokens', 'delete_denied_tokens',
               'create_roles', 'read_roles', 'update_roles', 'delete_roles',
               'create_users', 'read_users', 'update_users', 'delete_users'
    );
-- +goose StatementEnd
