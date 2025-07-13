
-- name: GetToDoItem :one
SELECT * FROM ToDoItems
WHERE Id = ? LIMIT 1;

-- name: ListToDoItems :many
SELECT * FROM ToDoItems;

-- name: InsertToDoItem :one
INSERT INTO ToDoItems (
  Id, Name, Description
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateToDoItem :one
UPDATE ToDoItems
SET 
  Name = ?,
  Description = ?
WHERE Id = ?
RETURNING *;

-- name: DeleteToDoItem :exec
DELETE FROM ToDoItems
WHERE Id = ?;