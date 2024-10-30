-- +goose Up
create table message (
    id serial primary key,
    chat_id int not null,
    sent_by text not null,
    text text not null,
    created_at timestamp not null default now(),
    updated_at timestamp,
    FOREIGN KEY (chat_id) REFERENCES chat(id)
    ON DELETE CASCADE
);

-- +goose Down
drop table message;

