-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_statistic (
                                           chat_id varchar(50) NOT NULL,
                                           test_id  integer,
                                           questions_number integer,
                                           correct_answers_number integer,
                                           FOREIGN KEY (test_id) REFERENCES tests(id),
                                           FOREIGN KEY (chat_id) REFERENCES "user"(chat_id),
                                           PRIMARY KEY(chat_id, test_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_statistic;
-- +goose StatementEnd
