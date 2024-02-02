package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerPostsFollowedByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limit := 10
	limit_s := r.URL.Query().Get("limit")
	if limit_s != "" {
		var err error
		limit, err = strconv.Atoi(limit_s)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "error while parsing query param")
			return
		}
	}

	postsFollowed, err := apiCfg.DB.GetPostsFollowedByUser(r.Context(), database.GetPostsFollowedByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("cannot find users followed feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusGone, dbPostsFollowedToPostsFollowed(postsFollowed))
}
