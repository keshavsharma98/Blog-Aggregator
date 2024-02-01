package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("Port not found int he environment")
	}

	chi_router := chi.NewRouter()

	chi_router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1_Router := chi.NewRouter()

	chi_router.Mount("/v1", v1_Router)

	v1_Router.Get("/readiness", handleReadiness)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: chi_router,
	}

	log.Printf("Server is running on port %v\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
