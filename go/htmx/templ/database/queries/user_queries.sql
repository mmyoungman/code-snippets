-- name: GetUser :one
SELECT * FROM Users
WHERE Id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM Users
ORDER BY Email;

-- name: InsertUser :one
INSERT INTO Users (
  Id, Username, Email, FirstName, LastName, RawIdToken
) VALUES (
  ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE Users
SET 
  Username = ?,
  Email = ?,
  FirstName = ?,
  LastName = ?,
  RawIdToken = ?
WHERE Id = ?;

-- name: DeleteUser :exec
DELETE FROM Users
WHERE Id = ?;