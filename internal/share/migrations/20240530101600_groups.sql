-- +goose Up
-- +goose StatementBegin
CREATE TABLE groups (
                            group_name varchar(255) NOT NULL PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE groups;
-- +goose StatementEnd
