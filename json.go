package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error occuring trying to marshal respone: %v", response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
