package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/manny-e1/rss_aggregator/internal/database"
)

type appConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT not found in the enviroment")
	}

	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("DB_URL not found in the enviroment")
	}

	conn, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatalf("Can't connect to database: %v", err)
	}

	app := appConfig{DB: database.New(conn)}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", checkHealth)
	v1Router.Post("/users", app.createUser)
	v1Router.Get("/users", app.authMiddleware(app.getUserByApiKey))
	v1Router.Post("/feeds", app.authMiddleware(app.createFeed))
	v1Router.Get("/feeds", app.getFeeds)
	v1Router.Post("/follow-feed", app.authMiddleware(app.createFeedFollow))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}
	log.Printf("Server started listening on %v", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
