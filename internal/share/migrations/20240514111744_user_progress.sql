-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_progress (
                               chat_id varchar(50) NOT NULL,
                               test_id integer NOT NULL,
                               question_id integer NOT NULL,
                               answer_id integer,
                               FOREIGN KEY (chat_id) REFERENCES "user"(chat_id),
                               FOREIGN KEY (answer_id) REFERENCES options(option_id),
                               FOREIGN KEY (test_id) REFERENCES tests(id),
                               FOREIGN KEY (question_id) REFERENCES questions(id),
                               PRIMARY KEY(chat_id, test_id, question_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_progress;
-- +goose StatementEnd
