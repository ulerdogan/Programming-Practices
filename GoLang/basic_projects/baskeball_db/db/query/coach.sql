-- name: CreateCoach :one
INSERT INTO coaches (
  username,
  hashed_password
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetCoachByUsername :one
SELECT * FROM coaches
WHERE username = $1 LIMIT 1;

-- name: GetCoachById :one
SELECT * FROM coaches
WHERE id = $1 LIMIT 1;

-- name: UpdateCoachPassword :one
UPDATE coaches
SET hashed_password = $2
WHERE username = $1
RETURNING *;

-- name: DeleteCoach :exec
DELETE FROM coaches
WHERE username = $1;