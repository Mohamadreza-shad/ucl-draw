-- name: InsertClub :one
INSERT INTO clubs (
  name,
  nationality,
  sedd
) VALUES (
  $1, $2, $3
)
RETURNING *;