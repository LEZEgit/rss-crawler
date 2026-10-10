-- +goose Up

CREATE TABLE feeds (
  id uuid primary key,
  created_at Timestamp not null,
  updated_at Timestamp not null,
  name Text not null,
  url TEXT unique not null,
  user_id uuid not null references users(id) on delete cascade
);

-- +goose Down

Drop table feeds;
