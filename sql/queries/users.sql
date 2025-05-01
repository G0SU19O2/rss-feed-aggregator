-- name: CreateUser :execresult
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    ?,
    ?,
    ?,
    ?
);
-- name: GetUser :one
SELECT * FROM users WHERE name = ? LIMIT 1;

-- name: DeleteUser :execresult
DELETE FROM users WHERE name = ? LIMIT 1;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;