package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"service/config"
	"service/handler"
	"service/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/lib/pq"
)

func main() {
	// database connection
	dbConfig := config.GetDbConfig()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully connnected to database.")
	defer db.Close()

	// instantiation dependencies
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", userHandler.RegisterUser)
			r.Post("/login", userHandler.Login)
		})
	})

	// start HTTP server
	log.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
