package app

import (
	"github.com/Dimoonevs/Online_store/internal/auth"
	conn "github.com/Dimoonevs/Online_store/internal/database/postgresql"
	"github.com/Dimoonevs/Online_store/internal/repository/postgresql"
	"github.com/Dimoonevs/Online_store/internal/service"
	"github.com/Dimoonevs/Online_store/internal/transport/rest"
	"github.com/Dimoonevs/Online_store/internal/transport/rest/handler"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func Run() {
	port := os.Getenv("PORT")
	secret := os.Getenv("SECRET")
	db := conn.ConnectToPostgres()

	log.Printf("Starting Service on port %s", port)

	// Set up repositories using PostgreSQL.
	repositories := postgresql.NewRepositories(db)

	// Initialize JWT (JSON Web Token) wrapper with provided secret key, issuer, and expiration time.
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       secret,
		Issuer:          "go-secrit-service",
		ExpirationHours: 24 * 100,
	}
	// Create a new service using the JWT wrapper and repositories.
	service := service.NewService(jwtWrapper, repositories)

	// Instantiate a handler with the service.
	handler := handler.NewHandler(service)

	// Create a new router for the REST API endpoints.
	router := rest.NewRouter(handler)

	// Create and start server on port 8080
	srv := &http.Server{
		Addr:    port,
		Handler: router.Router(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
