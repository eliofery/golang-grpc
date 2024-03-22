-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles_permissions (
    role_id integer NOT NULL,
    permission_id integer NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE roles_permissions;
-- +goose StatementEnd
