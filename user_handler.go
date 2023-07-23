package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/manny-e1/rss_aggregator/internal/auth"
	"github.com/manny-e1/rss_aggregator/internal/database"
)

func (app *appConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Params{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}
	user, err := app.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't create user: %v", err))
		return
	}
	respondWithJSON(w, 201, dbUserToCustomUser(user))
}

func (app *appConfig) getUserByApiKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("Auth error: %v", err))
		return
	}
	user, err := app.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 404, fmt.Sprint("user not found"))
		return
	}
	respondWithJSON(w, 201, dbUserToCustomUser(user))
}
