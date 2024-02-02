package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request JSON: %v", err))
		return
	}

	feedFollowed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("cannot create a new feed follow: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, dbFeedFollowToFeedFollow(feedFollowed))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	q_param := chi.URLParam(r, "feedFollowID")
	id, err := uuid.Parse(q_param)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing feedFollowID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     id,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("cannot delete the feed follow: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (apiCfg *apiConfig) handlerGetFeedsFollowedByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollowed, err := apiCfg.DB.GetFeedsFollowedByUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("cannot find users followed feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusGone, dbFeedsFollowToFeedsFollow(feedsFollowed))
}
