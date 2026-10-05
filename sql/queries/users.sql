-- name: CreateUser :one
Insert into users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
returning *;
