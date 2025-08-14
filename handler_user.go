package main

import (
	"net/http"
)

func handlerCreateUser(w http.ResponseWriter, r *http.Request){
	respondWithError(w, 200, "Something Went Wrong")
}