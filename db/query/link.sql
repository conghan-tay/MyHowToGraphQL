-- name: CreateLink :one
INSERT INTO links (
  title, address, user_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetLinks :many
SELECT links.id, links.title, links.address, links.user_id, users.username
FROM links INNER JOIN users
ON links.user_id = users.id;