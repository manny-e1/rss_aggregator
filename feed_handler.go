package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/manny-e1/rss_aggregator/internal/database"
)

func (app *appConfig) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type Params struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Params{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}
	feed, err := app.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't create feed: %v", err))
		return
	}
	respondWithJSON(w, 201, dbFeedToCustomFeed(feed))
}

func (app *appConfig) getFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := app.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't get feeds: %v", err))
		return
	}
	respondWithJSON(w, 201, dbFeedsToCustomFeeds(feeds))
}
