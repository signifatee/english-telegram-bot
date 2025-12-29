-- +goose Up
-- +goose StatementBegin
CREATE TABLE context (
    chat_id varchar(50) PRIMARY KEY NOT NULL,
    context  varchar(50) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE context;
-- +goose StatementEnd
