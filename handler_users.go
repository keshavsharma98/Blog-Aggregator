package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/keshavsharma98/Blog-Aggregator/internal/auth"
	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating new user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, dbUserToUser(user))
}

func (apiCfg *apiConfig) handleGetUserByApiKey(w http.ResponseWriter, r *http.Request) {
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

	respondWithJSON(w, http.StatusCreated, dbUserToUser(user))
}
