-- name: CreatePlayer :one
INSERT INTO players (
  name,
  role
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetPlayer :one
SELECT * FROM players
WHERE id = $1 LIMIT 1;

-- name: GetPlayerForUpdate :one
SELECT * FROM players
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: GetPlayers :many
SELECT * FROM players
WHERE name = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: GetTeamsPlayers :many
SELECT * FROM players
WHERE team = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdatePlayersTeam :one
UPDATE players
SET team = $2
WHERE id = $1
RETURNING *;

-- name: DeletePlayer :exec
DELETE FROM players
WHERE id = $1;