package handler

import (
	"github.com/google/uuid"
)

type parametersCreateUser struct {
	Name string `json:"name"`
}

type parametersCreateFeed struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type parametersCreateFeedFollow struct {
	FeedId uuid.UUID `json:"feed_id"`
}
