-- +goose Up
create table message (
    id serial primary key,
    text text not null,
    sent_by text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table message;

