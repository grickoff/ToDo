-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists tasks
(
    id SERIAL PRIMARY KEY,  -- Это id задачи, который будет пристраиваться к каждой новой. По нему можно обращаться.
    content VARCHAR(255) NOT NULL, -- описание задачи
    created_at timestamp default now() not null, -- время добавления
    deadline VARCHAR(255) NOT NULL, -- дата дедлайна
    status text NOT NULL-- статус задачи
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists users;
-- +goose StatementEnd

INSERT INTO tasks (content, created_at, deadline, status)
VALUES ('Научиться кодить, не доведя Серегу до инсульта', CURRENT_TIMESTAMP, '01.01.2024', 'Выполняется');