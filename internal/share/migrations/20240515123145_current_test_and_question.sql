-- +goose Up
-- +goose StatementBegin
CREATE TABLE current_test_and_question (
                         chat_id varchar(50) PRIMARY KEY  NOT NULL,
                         test_id  integer,
                         question_id  integer,
                         FOREIGN KEY (test_id) REFERENCES tests(id),
                         FOREIGN KEY (chat_id) REFERENCES "user"(chat_id),
                         FOREIGN KEY (question_id) REFERENCES questions(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE current_test_and_question;
-- +goose StatementEnd
