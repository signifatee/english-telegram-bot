-- +goose Up
-- +goose StatementBegin
CREATE TABLE questions (
                           id integer PRIMARY KEY UNIQUE ,
                           name varchar(255) NOT NULL,
                           right_option_id integer NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE questions;
-- +goose StatementEnd
