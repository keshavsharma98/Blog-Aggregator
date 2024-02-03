package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
)

// @Summary Create a new user
// @Description Creates a new user with the provided name
// @ID create-user
// @Accept json
// @Produce json
// @Param request body parametersCreateUser true "User creation parameters"
// @Success 201 {object} User
// @Failure 400 {object} errorResponseBody
// @Router /users [post]
func (apiCfg *ApiConfig) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := parametersCreateUser{}
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

// @Summary Get user by API key
// @Description Retrieves user details based on the provided API key
// @ID get-user-by-api-key
// @Produce json
// @Param Authorization header string true "API key for authentication"
// @Success 200 {object} User
// @Failure 401 {object} errorResponseBody
// @Router /users [get]
func (apiCfg *ApiConfig) HandleGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusCreated, dbUserToUser(user))
}
