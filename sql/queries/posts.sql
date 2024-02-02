-- name: CreatePosts :exec
INSERT INTO posts (id, title, url, description, published_at, feed_id, updated_at, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetPostsFollowedByUser :many
SELECT p.id,
p.title,
p.url,
p.description,
p.feed_id,
p.published_at,
p.created_at,
p.updated_at
FROM posts AS p
INNER JOIN feed AS f
ON p.feed_id=f.id
WHERE f.user_id =($1)
ORDER BY p.published_at DESC
LIMIT ($2); 
