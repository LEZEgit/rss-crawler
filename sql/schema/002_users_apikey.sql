-- +goose Up

ALTER table users
  add column api_key varchar(64) UNIQUE not null default gen_random_uuid()::varchar(64);

-- +goose Down
ALTER table users
  drop column api_key;
