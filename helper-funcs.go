package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponseBody struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	body := errorResponseBody{
		Error: msg,
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	marshaled_payload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal the response: %v\n", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(marshaled_payload)
}
