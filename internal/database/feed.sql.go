// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: feed.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feed (id, name, url, user_id, updated_at,created_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, name, url, user_id, created_at, updated_at
`

type CreateFeedParams struct {
	ID        uuid.UUID
	Name      string
	Url       string
	UserID    uuid.UUID
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.Name,
		arg.Url,
		arg.UserID,
		arg.UpdatedAt,
		arg.CreatedAt,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
