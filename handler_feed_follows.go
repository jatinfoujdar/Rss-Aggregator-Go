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

    feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
        ID:        uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        UserID:    user.ID,
        FeedID:   params.FeedId,
    })
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
        return
    }

    respondWithJson(w, 201, DatabaseFeedFollowToFeedFollow(feedFollow))
}



