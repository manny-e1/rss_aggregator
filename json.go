package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error occuring trying to marshal respone: %v", response)
		return
	}
	fmt.Println(data)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
