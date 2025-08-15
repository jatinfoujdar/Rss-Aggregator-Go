package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/google/uuid"
    "github.com/jatinfoujdar/Rss-Aggregator-Go/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
    type parameters struct {
        Name string `json:"name"`
    }
    decoder := json.NewDecoder(r.Body)
    params := parameters{}
    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
        return
    }

    user := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
        ID:        uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        Name:      params.Name,
    })

respondWithJson(w, 200, user) 
}