package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type NewFeedToFeed struct {
	Feed       Feed       `json:"feed"`
	FeedFollow FeedFollow `json:"feed_follow"`
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

func dbFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url,
		UserID:    dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func dbNewFeedToFeed(dbFeed database.Feed, dbFeedFollowed database.FeedFollow) NewFeedToFeed {
	return NewFeedToFeed{
		Feed: Feed{
			ID:        dbFeed.ID,
			Name:      dbFeed.Name,
			URL:       dbFeed.Url,
			UserID:    dbFeed.UserID,
			CreatedAt: dbFeed.CreatedAt,
			UpdatedAt: dbFeed.UpdatedAt,
		},
		FeedFollow: dbFeedFollowToFeedFollow(dbFeedFollowed),
	}
}

func dbFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feedsArr := make([]Feed, 0, len(dbFeeds))
	for _, df := range dbFeeds {
		f := dbFeedToFeed(df)
		feedsArr = append(feedsArr, f)
	}
	return feedsArr
}

func dbFeedFollowToFeedFollow(dbFeedsFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedsFollow.FeedID,
		UserID:    dbFeedsFollow.UserID,
		FeedID:    dbFeedsFollow.FeedID,
		CreatedAt: dbFeedsFollow.CreatedAt,
		UpdatedAt: dbFeedsFollow.UpdatedAt,
	}
}

func dbFeedsFollowToFeedsFollow(dbFeedsFollow []database.FeedFollow) []FeedFollow {
	feedsArr := make([]FeedFollow, 0, len(dbFeedsFollow))
	for _, df := range dbFeedsFollow {
		f := dbFeedFollowToFeedFollow(df)
		feedsArr = append(feedsArr, f)
	}
	return feedsArr
}
