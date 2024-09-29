-- name: InsertClub :one
INSERT INTO clubs (
  name,
  nationality,
  seed
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAllClubs :many
SELECT * FROM clubs;