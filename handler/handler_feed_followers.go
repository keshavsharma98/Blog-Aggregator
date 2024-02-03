package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
)

// handlerCreateFeedFollow creates a new feed follow.
// @Summary Create a new feed follow
// @Description Creates a new feed follow for the user
// @ID create-feed-follow
// @Accept json
// @Produce json
// @Param request body parametersCreateFeedFollow true "Feed follow parameters"
// @Success 201 {object} FeedFollow
// @Failure 400 {object} errorResponseBody
// @Router /feed-follows [post]
func (apiCfg *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := parametersCreateFeedFollow{}
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

// handlerDeleteFeedFollow deletes a feed follow.
// @Summary Delete a feed follow
// @Description Deletes a feed follow for the user
// @ID delete-feed-follow
// @Param feedFollowID path string true "Feed follow ID"
// @Success 200
// @Failure 400 {object} errorResponseBody
// @Router /feed-follows/{feedFollowID} [delete]
func (apiCfg *ApiConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
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

// handlerGetFeedsFollowedByUser gets feeds followed by the user.
// @Summary Get feeds followed by the user
// @Description Retrieves a list of feeds followed by the user
// @ID get-feeds-followed-by-user
// @Produce json
// @Success 200 {array} FeedFollow
// @Failure 400 {object} errorResponseBody
// @Router /feed-followed [get]
func (apiCfg *ApiConfig) HandlerGetFeedsFollowedByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollowed, err := apiCfg.DB.GetFeedsFollowedByUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("cannot find users followed feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusGone, dbFeedsFollowToFeedsFollow(feedsFollowed))
}
