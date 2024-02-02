package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

var n = 10

var feedsUrls = []string{
	"https://blog.boot.dev/index.xml",
	"https://wagslane.dev/index.xml",
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

	defer conn.Close()

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
	v1_Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUserByApiKey))
	v1_Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1_Router.Get("/feed", apiCfg.handlerGetAllFeeds)
	v1_Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1_Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))
	v1_Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedsFollowedByUser))

	chi_router.Mount("/v1", v1_Router)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: chi_router,
	}

	rssCrawler()

	log.Printf("Server is running on port %v\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}

func rssCrawler() {
	var wg sync.WaitGroup

	for _, u := range feedsUrls {

		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			j := 0
			for {
				j = Crawler(u, n, j)
				j = j + n
				time.Sleep(time.Second * 5)
			}
		}(u)
	}
	wg.Wait()
}
