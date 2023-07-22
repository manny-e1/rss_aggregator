package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT not found in the enviroment")
	}

	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}
	log.Printf("Server started listening on %v", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
