-- +goose Up

CREATE TABLE users (
  id uuid primary key,
  created_at Timestamp not null,
  updated_at Timestamp not null,
  name Text not null
);

-- +goose Down

Drop table users;
