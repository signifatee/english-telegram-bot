-- +goose Up
-- +goose StatementBegin
CREATE TABLE options (
                         option_id integer NOT NULL UNIQUE,
                         name varchar(255) NOT NULL,
                         question_id integer NOT NULL,
                         FOREIGN KEY (question_id) REFERENCES questions(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE options;
-- +goose StatementEnd
