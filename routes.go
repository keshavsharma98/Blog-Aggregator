package main

import (
	"github.com/go-chi/chi"
	"github.com/keshavsharma98/Blog-Aggregator/handler"
)

func getV1Routes(apiCfg *handler.ApiConfig) *chi.Mux {
	v1_Router := chi.NewRouter()

	// users routes
	v1_Router.Post("/users", apiCfg.HandleCreateUser)
	v1_Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandleGetUserByApiKey))

	//feeds routes
	v1_Router.Post("/feed", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	v1_Router.Get("/feed", apiCfg.HandlerGetAllFeeds)

	// feeds follow routes
	v1_Router.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	v1_Router.Delete("/feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollow))
	v1_Router.Get("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedsFollowedByUser))

	//posts routes
	v1_Router.Get("/posts", apiCfg.MiddlewareAuth(apiCfg.HandlerPostsFollowedByUser))

	return v1_Router
}
