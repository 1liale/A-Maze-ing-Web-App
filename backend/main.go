// Package main is the root package managing the gin-gonic REST API.

package main

import (
	"io"
	"os"

	"github.com/1liale/maze-backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

var logger *logrus.Logger

func init() {
	// load env
	err := godotenv.Load()
	if err != nil {
	}

	// configure logger
	logger = logrus.New()
	logLevel, err := logrus.ParseLevel(os.Getenv("LOGGER_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	f, _ := os.Create("log.out")
	channels := io.MultiWriter(f, os.Stdout)

	logger.SetLevel(logLevel)
	logger.SetOutput(channels)
	logger.SetReportCaller(true)
}

func main() {
	// init gin server
	port := ":8080"
	router := gin.New()

	// middlewares
	router.Use(
		ginlogrus.Logger((logger)),
		gin.Recovery(),
	)

	// group CRUD endpoints
	v1 := router.Group("/api/v1")
	{
		v1.GET("/api-health", handlers.SystemCheck)
	}

	router.GET("/api-health", handlers.SystemCheck)
	router.Run(port)
}
