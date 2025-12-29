-- +goose Up
-- +goose StatementBegin
CREATE TABLE available_tests_for_user (
                                          chat_id varchar(10),
                                          test_id integer NOT NULL,
                                          FOREIGN KEY (test_id) REFERENCES tests(id),
                                          FOREIGN KEY (chat_id) REFERENCES "user"(chat_id),
                                          PRIMARY KEY(chat_id, test_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE available_tests_for_user;
-- +goose StatementEnd
