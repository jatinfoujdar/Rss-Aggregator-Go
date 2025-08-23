package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    

    "github.com/google/uuid"
    "github.com/jatinfoujdar/Rss-Aggregator-Go/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
    type parameters struct {
            FeedId uuid.UUID `json:"feed_id"`
    }
    decoder := json.NewDecoder(r.Body)
    params := parameters{}
    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
        return
    }

    feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
        ID:        uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        Name:      params.Name,
        Url:       params.URL,
        UserID:    user.ID,
    })
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
        return
    }

    respondWithJson(w, 201, DatabaseFeedToFeed(feed))
}



