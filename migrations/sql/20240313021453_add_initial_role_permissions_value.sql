-- +goose Up
-- +goose StatementBegin
INSERT INTO role_permissions (role_id, permission_id)
VALUES (1, 1), -- admin (all permissions)
       (1, 2),
       (1, 3),
       (1, 4),
       (1, 5),
       (1, 6),
       (1, 7),
       (1, 8),
       (1, 9),
       (1, 10),
       (1, 11),
       (1, 12),
       (1, 13),
       (1, 14),
       (1, 15),
       (1, 16),
       (1, 17),
       (1, 18),
       (1, 19),
       (1, 20),

       (2, 1), -- moderator (except delete)
       (2, 2),
       (2, 3),
       (2, 5),
       (2, 6),
       (2, 7),
       (2, 9),
       (2, 10),
       (2, 11),
       (2, 13),
       (2, 14),
       (2, 15),
       (2, 17),
       (2, 18),
       (2, 19),

       (3, 2), -- user (only read)
       (3, 6),
       (3, 10),
       (3, 14),
       (3, 18);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM role_permissions
WHERE role_id IN (1,2,3);
-- +goose StatementEnd
