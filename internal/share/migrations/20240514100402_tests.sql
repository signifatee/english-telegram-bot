-- +goose Up
-- +goose StatementBegin
CREATE TABLE tests (
                            id integer PRIMARY KEY NOT NULL UNIQUE,
                            name varchar(150) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tests;
-- +goose StatementEnd
