package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
)

// handlerCreateFeed creates a new feed.
// @Summary Create a new feed
// @Description Creates a new feed with the provided name and URL
// @ID create-feed
// @Accept json
// @Produce json
// @Param request body parametersCreateFeed true "Feed creation parameters"
// @Success 201 {object} Feed
// @Failure 400 {object} errorResponseBody
// @Router /feed [post]
func (apiCfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := parametersCreateFeed{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:            uuid.New(),
		Name:          params.Name,
		Url:           params.URL,
		UserID:        user.ID,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		LastFetchedAt: sql.NullTime{},
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("cannot create a new feed: %v", err))
		return
	}

	feedFollowed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("cannot create a new feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, dbNewFeedToFeed(feed, feedFollowed))
}

// handlerGetAllFeeds gets all feeds.
// @Summary Get all feeds
// @Description Retrieves a list of all feeds
// @ID get-all-feeds
// @Produce json
// @Success 200 {array} Feed
// @Failure 400 {object} errorResponseBody
// @Router /feed [get]
func (apiCfg *ApiConfig) HandlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("cannot get feeds: %v", err))
		return
	}

	respondWithJSON(w, http.StatusFound, dbFeedsToFeeds(feeds))
}
