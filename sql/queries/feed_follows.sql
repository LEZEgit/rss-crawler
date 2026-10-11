-- name: CreateFeedFollow :one
Insert into feed_follows (id, created_at, updated_at, user_id, feed_id)
Values ($1, $2, $3, $4, $5)
Returning *;

-- name: GetFeedsFollowingByUserID :many
select f.* from feeds f
join feed_follows ff on f.id = ff.feed_id
where ff.user_id = $1;

-- name: UnfollowFeed :exec
Delete from feed_follows
where user_id = $1 and feed_id = $2;