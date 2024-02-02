-- name: CreateFeed :one
INSERT INTO feed (id, name, url, user_id, updated_at,created_at,last_fetched_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feed;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feed
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT ($1);

-- name: MarkFetched :exec
UPDATE feed
SET updated_at=($1), last_fetched_at=($2)
WHERE id=($3);