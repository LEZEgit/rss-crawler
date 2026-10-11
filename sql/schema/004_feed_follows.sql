-- +goose Up
Create table feed_follows (
  id uuid primary key,
  created_at Timestamp not null,
  updated_at Timestamp not null,
  user_id uuid not null references users(id) on delete cascade,
  feed_id uuid not null references feeds(id) on delete cascade
);
Create unique index feed_follows_user_id_feed_id_idx on feed_follows(user_id, feed_id);

-- +goose Down  
drop table feed_follows;
drop index feed_follows_user_id_feed_id_idx;