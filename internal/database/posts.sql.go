// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: posts.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPosts = `-- name: CreatePosts :exec
INSERT INTO posts (id, title, url, description, published_at, feed_id, updated_at, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type CreatePostsParams struct {
	ID          uuid.UUID
	Title       string
	Url         string
	Description *string
	PublishedAt time.Time
	FeedID      uuid.UUID
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

func (q *Queries) CreatePosts(ctx context.Context, arg CreatePostsParams) error {
	_, err := q.db.ExecContext(ctx, createPosts,
		arg.ID,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
		arg.UpdatedAt,
		arg.CreatedAt,
	)
	return err
}

const getPostsFollowedByUser = `-- name: GetPostsFollowedByUser :many
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
LIMIT ($2)
`

type GetPostsFollowedByUserParams struct {
	UserID uuid.UUID
	Limit  int32
}

func (q *Queries) GetPostsFollowedByUser(ctx context.Context, arg GetPostsFollowedByUserParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsFollowedByUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.FeedID,
			&i.PublishedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
