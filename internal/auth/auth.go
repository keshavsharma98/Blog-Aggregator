package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(h http.Header) (string, error) {
	api_key := h.Get("Authorization")
	if api_key == "" {
		return "", errors.New("no authentication info found")
	}

	key := strings.Split(api_key, " ")
	if len(key) != 2 {
		return "", errors.New("no authentication info found")
	}

	if key[0] != "ApiKey" {
		return "", errors.New("malformed auth header")
	}

	return key[1], nil
}
