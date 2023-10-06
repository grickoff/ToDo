-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists goose_db_version
(
    id         serial  not null
    constraint goose_db_version_pkey
    primary key,
    version_id bigint  not null,
    is_applied boolean not null,
    tstamp     timestamp default now()
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
