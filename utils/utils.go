package utils

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, data interface{})  {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}