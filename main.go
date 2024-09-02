package main

import (
	"auth-service/internal/handlers"
	"auth-service/internal/repositories"
	"auth-service/internal/services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"log"
)

func main() {
	router := gin.Default()

	connPool, err := initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v\n", err)
	}
	defer connPool.Close()

	logger := initLogger()
	defer logger.Sync()

	initAuthHandlers(router, logger, connPool)

	router.Run("localhost:8080")
}

func initAuthHandlers(router *gin.Engine, logger *zap.SugaredLogger, connPool *pgxpool.Pool) {
	var userRepository = repositories.NewUserRepositoryImpl(connPool, logger)
	var userService = services.NewUserService(userRepository, logger)
	var regHandler = handlers.NewUserRegHandler(userService, logger)
	var authHandler = handlers.NewUserAuthHandler(userService, logger)

	router.POST("/api/v1/register", regHandler.Register)
	router.POST("/api/v1/authenticate", authHandler.Authenticate)
}

func initLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Exception during logger initizalization")
	}
	return logger.Sugar()
}

func initDB() (*pgxpool.Pool, error) {
	databaseUrl := "postgres://anabiev:qwerty007@localhost:5436/awesomeProject"

	connPool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	err = connPool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to ping the database: %v", err)
	}

	return connPool, nil
}
