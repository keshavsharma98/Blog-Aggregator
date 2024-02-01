package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	var (
		dbURL string
		port  string
	)

	port = os.Getenv("PORT")
	if port == "" {
		log.Panicln("Port not found")
	}

	dbURL = os.Getenv("POSTGRES_URL")
	if dbURL == "" {
		log.Panicln("Database URL not found")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Panicln("Error connecting to database ", err)
	}

	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
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
	chi_router.Get("/readiness", apiCfg.handleReadiness)
	v1_Router.Post("/users", apiCfg.handleCreateUser)
	v1_Router.Get("/users", apiCfg.handleGetUserByApiKey)

	chi_router.Mount("/v1", v1_Router)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: chi_router,
	}

	log.Printf("Server is running on port %v\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
