package auth

import (
	"errors"
	"net/http"
	"strings"
)
func GetAPIKey(headers http.Header)(string, error){
	val := headers.Get("Authorization")
	if val == ""{
		return "", errors.New("api key not found in headers")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("invalid API key format")
	}
	return vals[1], nil
}