-- +goose Up
-- +goose StatementBegin
CREATE TABLE egts_history_records (
   id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
   egts_package_hex TEXT UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
