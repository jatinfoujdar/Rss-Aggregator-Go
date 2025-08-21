package main

import (
	"fmt"
	"net/http"

	"github.com/jatinfoujdar/Rss-Aggregator-Go/internal/auth"
	"github.com/jatinfoujdar/Rss-Aggregator-Go/internal/database"
)


type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	apiKey, err := auth.GetAPIKey(r.Header)
    if err != nil {
        respondWithError(w, http.StatusForbidden, fmt.Sprintf("Error getting API key: %v", err))
        return
    }

    user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
    if err != nil {
        respondWithError(w, http.StatusNotFound, fmt.Sprintf("User not found: %v", err))
        return
    }

		handler(w, r, user)
	}
}