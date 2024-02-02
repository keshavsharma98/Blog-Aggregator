-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, user_id, feed_id, updated_at,created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE id=($1);

-- name: GetFeedsFollowedByUser :many
SELECT * FROM feed_follows
WHERE user_id=($1);

