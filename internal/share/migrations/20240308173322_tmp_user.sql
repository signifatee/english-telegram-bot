-- +goose Up
-- +goose StatementBegin
CREATE TABLE "tmp_user" (
                        chat_id varchar(50) PRIMARY KEY NOT NULL,
                        name  varchar(255),
                        institute varchar(50),
                        "group" varchar(50),
                        language_level varchar(2)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "tmp_user";
-- +goose StatementEnd
