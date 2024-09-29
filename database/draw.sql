-- name: CreateMatch :one
INSERT INTO draws (
  host_id,
  guest_id
) VALUES (
  $1, $2
)
RETURNING *;