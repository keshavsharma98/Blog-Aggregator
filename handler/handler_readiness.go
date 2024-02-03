package handler

import "net/http"

// handleReadiness checks the readiness of the server.
// @Summary Check server readiness
// @Description Checks if the server is ready to handle requests
// @ID check-readiness
// @Produce plain
// @Success 200 {string} string "Server is running"
// @Failure 500 {object} errorResponseBody "Internal Server Error"
// @Router /readiness [get]
func (apiCfg *ApiConfig) HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "Server is running")
}
