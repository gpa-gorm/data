package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// MergeRouters merges multiple routers into a single router
func MergeRouters(baseRouter chi.Router, basePath string, routers ...*chi.Mux) {
	for _, router := range routers {
		baseRouter.Mount(basePath, router)
	}
}

// InitBaseRouter initializes the base router with common middlewares
func InitBaseRouter(corsAllowedOrigins []string) *chi.Mux {

	// Define the base router for sharing with subrouters
	// This approach is used to avoid the need to pass the base router to each subrouter
	// Also it provides use common middlewares for all subrouters and specific middlewares for each subrouter
	r := chi.NewRouter()

	// Define the middlewares for the base router
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   corsAllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	return r
}
