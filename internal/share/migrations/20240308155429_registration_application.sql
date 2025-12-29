-- +goose Up
-- +goose StatementBegin
CREATE TABLE registration_application (
                      chat_id varchar(50) PRIMARY KEY NOT NULL,
                      status  varchar(50) NOT NULL,
                      CONSTRAINT unique_chat_id UNIQUE (chat_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE registration_application;
-- +goose StatementEnd
