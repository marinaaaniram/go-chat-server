-- +goose Up
ALTER TABLE public.chat
DROP COLUMN usernames;

-- +goose Down
ALTER TABLE public.chat
ADD COLUMN usernames text[];
