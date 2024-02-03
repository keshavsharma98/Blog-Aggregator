package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/keshavsharma98/Blog-Aggregator/docs"
	"github.com/keshavsharma98/Blog-Aggregator/handler"
	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
	"github.com/keshavsharma98/Blog-Aggregator/internal/scrapper"
	_ "github.com/lib/pq"
	"github.com/swaggo/http-swagger"
)

// @title Orders API
// @version 1.0
// @description This is a sample service for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /v1

func main() {
	godotenv.Load(".env")
	var port = os.Getenv("PORT")
	if port == "" {
		log.Panicln("Port not found")
	}

	queries, conn := newDB()
	apiCfg := handler.ApiConfig{
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

	chi_router.Get("/readiness", apiCfg.HandlerReadiness)

	v1_Router := getV1Routes(&apiCfg)
	chi_router.Mount("/v1", v1_Router)

	chi_router.Route("/swagger", func(r chi.Router) {
		r.Get("/*", httpSwagger.WrapHandler)
	})

	server := &http.Server{
		Addr:    ":" + port,
		Handler: chi_router,
	}

	scrappingOfFeeds(conn)
	defer conn.Close()

	log.Printf("Server is running on port %v\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}

func newDB() (*database.Queries, *sql.DB) {
	var dbURL = os.Getenv("POSTGRES_URL")
	if dbURL == "" {
		log.Panicln("Database URL not found")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Panicln("Error connecting to database ", err)
	}

	queries := database.New(conn)
	return queries, conn
}

func scrappingOfFeeds(conn *sql.DB) {
	db := database.New(conn)
	concurrency_s := os.Getenv("CONCURRENCY")
	if concurrency_s == "" {
		log.Panicln("concurrency number not found")
	}

	duration_s := os.Getenv("CONCURRENCY_WAIT")
	if duration_s == "" {
		log.Panicln("concurrency wait duration not found")
	}
	concurrency, err := strconv.Atoi(concurrency_s)
	if err != nil {
		log.Panicln("Error: ", err)
	}
	duration, err := time.ParseDuration(duration_s)
	if err != nil {
		log.Panicln("Error: ", err)
	}

	runScrapping := os.Getenv("RUN_SCRAPPING")
	if runScrapping == "" || runScrapping == "true" {
		scrapper.RssScraper(db, concurrency, duration)
	}

}
