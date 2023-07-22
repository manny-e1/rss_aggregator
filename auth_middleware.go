package main

import (
	"fmt"
	"net/http"

	"github.com/manny-e1/rss_aggregator/internal/auth"
	"github.com/manny-e1/rss_aggregator/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (app *appConfig) authMiddleware(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := app.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, "user not found")
			return
		}
		handler(w, r, user)
	}
}
