-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles
(
    id   serial primary key,
    name varchar(256) not null unique
);

INSERT INTO roles (name) VALUES ('user') ON CONFLICT (name) DO NOTHING;
INSERT INTO roles (name) VALUES ('admin') ON CONFLICT (name) DO NOTHING;

CREATE TABLE users
(
    id         serial primary key,
    name       varchar(256) not null,
    email      varchar(256) not null unique,
    pass_hash  varchar,
    role       int references roles (id) on delete cascade not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
drop table roles;
-- +goose StatementEnd