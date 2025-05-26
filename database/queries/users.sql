-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (id, username, email, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
  username = COALESCE(sqlc.narg('username'), username),
  password = COALESCE(sqlc.narg('password'), password),
  email = COALESCE(sqlc.narg('email'), email),
  updated_at = now()
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;
