package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "messageOK/docs"

	_ "github.com/go-sql-driver/mysql"

	"messageOK/config"
	"messageOK/internal/handler"
	"messageOK/internal/repository"
	"messageOK/internal/usecase"
	"messageOK/pkg/redis"
)

func main() {
	// Loading configuration
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Getting environmental values and setting up My sql connection
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = cfg.MySQL.Host
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = cfg.MySQL.Port
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = cfg.MySQL.User
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = cfg.MySQL.Password
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = cfg.MySQL.DBName
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the mysql database:", err)
	}

	// Getting environmental values and Setup Redis connection
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = cfg.Redis.Host
	}
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = cfg.Redis.Port
	}

	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	rc := redis.NewRedisClient(redisAddr)

	// Dependency Injection
	repo := repository.NewMySQLMessageRepository(db)
	useCase := usecase.NewMessageUseCase(repo, rc, cfg)

	router := mux.NewRouter()
	handler.NewMessageHandler(router, useCase)

	// Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// CORS configuration
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "X-CSRF-Token"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	// Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server at port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(router)))
}
