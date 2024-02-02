package main

import (
	"database/sql"
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

type Post struct {
	ID          uuid.UUID      `json:"id"`
	Title       string         `json:"title"`
	Url         string         `json:"url"`
	Description sql.NullString `json:"description"`
	FeedID      uuid.UUID      `json:"feed_id"`
	PublishedAt time.Time      `json:"published_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
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

func dbPostToPost(postsFollowed database.Post) Post {
	return Post{
		ID:          postsFollowed.ID,
		Title:       postsFollowed.Title,
		Url:         postsFollowed.Url,
		Description: postsFollowed.Description,
		FeedID:      postsFollowed.FeedID,
		PublishedAt: postsFollowed.PublishedAt,
		CreatedAt:   postsFollowed.CreatedAt,
		UpdatedAt:   postsFollowed.UpdatedAt,
	}
}

func dbPostsFollowedToPostsFollowed(postsFollowed []database.Post) []Post {
	postsArr := make([]Post, 0, len(postsFollowed))
	for _, pf := range postsFollowed {
		p := dbPostToPost(pf)
		postsArr = append(postsArr, p)
	}
	return postsArr
}
