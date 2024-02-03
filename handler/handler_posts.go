package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
)

// handlerPostsFollowedByUser gets posts followed by the user.
// @Summary Get posts followed by the user
// @Description Retrieves a list of posts followed by the user
// @ID get-posts-followed-by-user
// @Param limit query int false "Limit the number of posts (default: 10)"
// @Produce json
// @Success 200 "with an array of Posts"
// @Failure 400 {object} errorResponseBody
// @Router /posts [get]
func (apiCfg *ApiConfig) HandlerPostsFollowedByUser(w http.ResponseWriter, r *http.Request, user database.User) {
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

	respondWithJSON(w, http.StatusOK, dbPostsFollowedToPostsFollowed(postsFollowed))
}
