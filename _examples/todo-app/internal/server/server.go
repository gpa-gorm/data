package server

import (
	"fmt"
	"net/http"
	"time"

	"gpadata/todoapp/config"
	"gpadata/todoapp/internal/db"
	"gpadata/todoapp/internal/routes"

	todo "gpadata/todoapp/internal/app/todo"
)

// RunServer starts the server
func RunServer() {


	// Validate and Initialize configs from .env file
	config.Validate()

	// Set Timezone
	loc, _ := time.LoadLocation("Europe/Istanbul")
	time.Local = loc

	// Postgres
	postgresDSN, _ := config.DbConfiguration(&config.DBAppCfg)
	if err := db.InitPostgres(postgresDSN); err != nil {
		fmt.Printf("database DbConnection error: %s", err)
	}

	// Migration
	db.InitialMigration(db.GetPostgresDB())

	// Initialize Base Router
	handler := routes.InitBaseRouter([]string{"*"})

	// Initialize repositories
	todoRepository := todo.NewTodoRepository(db.GetPostgresDB()).GetInstance()

	// Initialize services
	todoService := todo.NewTodoService(todoRepository).GetInstance()

	// Initialize controllers
	todoController := todo.NewTodoController(todoService).GetInstance()

	// Initialize routes
	todoRoutes := todo.NewTodoRoutes(todoController)

	// Merge base routes with subroutes
	routes.MergeRouters(handler, "/api/v1/todo", todoRoutes.GetRouter())

	// Initialize server
	serve := &http.Server{
		Addr:         ":8090",
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// logger.Println(fmt.Sprintf("Starting server on port %d", cfg.Port))
	fmt.Println("Starting server on port 8090...")

	// Start server
	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
