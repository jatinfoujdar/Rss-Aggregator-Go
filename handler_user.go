package main

import (
	"encoding/json"
	"fmt"
	
	"net/http"

	"github.com/google/uuid"
	"github.com/jatinfoujdar/Rss-Aggregator-Go/internal/database"
)

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request){

type parameters struct{
	Name string `json:"name"`
}
decoder :=  json.NewDecoder(r.Body)
params := parameters{}
// decoder.Decode(&params)
err := decoder.Decode(&params)
if err != nil {
	respondWithError(w,400, fmt.Sprintf("Error parsing JSON: %v", err))
return 
}

apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
	ID: uuid.New(),
	Name: params.Name,

})
	respondWithJson(w,200, struct{}{})
}