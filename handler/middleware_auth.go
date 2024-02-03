package handler

import (
	"fmt"
	"net/http"

	"github.com/keshavsharma98/Blog-Aggregator/internal/auth"
	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) MiddlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api_key, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Authentication error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), api_key)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Cannot find user: %v", err))
			return
		}

		handler(w, r, user)
	}

}
