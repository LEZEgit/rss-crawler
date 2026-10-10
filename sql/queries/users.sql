-- name: CreateUser :one
Insert into users (id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4,
    gen_random_uuid()::varchar(64)
)
returning *;

-- name: GetUserByAPIKey :one
Select * from users where api_key = $1;