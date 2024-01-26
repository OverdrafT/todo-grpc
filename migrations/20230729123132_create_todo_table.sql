-- +goose Up
create table todo (
    id serial primary key,
    title text not null,
    content text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table todo;

