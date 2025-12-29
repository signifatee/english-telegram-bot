-- +goose Up
-- +goose StatementBegin
CREATE TABLE institutes (
                            institute_name varchar(50) NOT NULL PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE institutes;
-- +goose StatementEnd
