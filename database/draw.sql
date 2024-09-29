-- name: CreateMatch :one
INSERT INTO draws (
  host_id,
  guest_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetMatches :many
SELECT 
d.id, 
host.name as host_name, 
away.name as away_name 
FROM draws as d
Join clubs as host on d.host_id = host.id
Join clubs as away on d.guest_id = away.id;