-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
                      chat_id varchar(50) PRIMARY KEY  NOT NULL,
                      name  varchar(255) NOT NULL,
                      institute varchar(50) NOT NULL,
                      "group" varchar(50) NOT NULL,
                      language_level varchar(2) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user";
-- +goose StatementEnd
