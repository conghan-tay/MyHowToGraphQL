-- name: CreateUser :one
INSERT INTO users (
  username, password
) VALUES (
  $1, $2
)
RETURNING id, username;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;