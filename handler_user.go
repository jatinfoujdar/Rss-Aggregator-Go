package main

import (
	"net/http"
	
)

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request){
	respondWithJson(w,200, struct{}{})
}