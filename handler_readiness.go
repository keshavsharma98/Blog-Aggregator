package main

import "net/http"

func (apiCfg *apiConfig) handleReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "Server is running")
}
