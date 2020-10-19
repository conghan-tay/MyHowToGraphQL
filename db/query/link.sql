-- name: CreateLink :one
INSERT INTO links (
  title, address
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetLinks :many
SELECT * FROM links;