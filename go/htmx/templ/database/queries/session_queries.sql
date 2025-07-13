-- name: GetSession :one
SELECT * FROM Sessions
WHERE Id = ?
LIMIT 1;

-- name: InsertSession :one
INSERT INTO Sessions (
  Id, UserId, AccessToken, RefreshToken, Expiry, TokenType
) VALUES (
  ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateSession :exec
UPDATE Sessions
SET 
  AccessToken = ?,
  RefreshToken = ?,
  Expiry = ?,
  TokenType = ?
WHERE Id = ?
AND UserId = ?;

-- name: DeleteSession :exec
DELETE FROM Sessions
WHERE Id = ?;