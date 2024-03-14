-- +goose Up
-- +goose StatementBegin
CREATE TABLE role_permissions (
    role_id integer NOT NULL,
    permission_id integer NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);

CREATE INDEX idx_role_permissions_role_id ON role_permissions (role_id);
CREATE INDEX idx_role_permissions_permission_id ON role_permissions (permission_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE role_permissions;

DROP INDEX idx_role_permissions_role_id;
DROP INDEX idx_role_permissions_permission_id;
-- +goose StatementEnd
