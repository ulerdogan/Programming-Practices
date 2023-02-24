-- name: CreateTeam :one
INSERT INTO teams (
  coach_id
) VALUES (
  $1
) RETURNING *;

-- name: GetTeam :one
SELECT * FROM teams
WHERE id = $1 LIMIT 1;

-- name: GetTeamByCoach :one
SELECT * FROM teams
WHERE coach_id = $1 LIMIT 1;

-- name: ListTeams :many
SELECT * FROM teams
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTeamsWin :one
UPDATE teams
SET wins = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTeam :exec
DELETE FROM teams
WHERE id = $1;

-- name: IncreaseWinTeam :one
UPDATE teams
SET wins = wins + $2
WHERE id = $1
RETURNING *;
