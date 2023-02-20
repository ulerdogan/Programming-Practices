-- name: CreateTransfer :one
INSERT INTO transfers (
  player,
  from_team,
  to_team
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
WHERE 
    from_team = $1 OR
    to_team = $2
ORDER BY id
LIMIT $3
OFFSET $4;
