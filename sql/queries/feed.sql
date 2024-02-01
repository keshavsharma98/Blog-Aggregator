-- name: CreateFeed :one
INSERT INTO feed (id, name, url, user_id, updated_at,created_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feed;
