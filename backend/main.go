// Package main is the root package managing the gin-gonic REST API.

package main

import (
	"fmt"

	"github.com/1liale/maze-backend/handlers"
	"github.com/1liale/maze-backend/middlewares"

	"github.com/gin-gonic/gin"

	ginlogrus "github.com/toorop/gin-logrus"
)

var logger *middlewares.CustomStdLogger

func init() {
	logger = middlewares.NewLogger()
}

func main() {
	// init gin server
	port := ":8080"
	router := gin.New()

	// middlewares
	router.Use(
		ginlogrus.Logger(logger),
		gin.Recovery(),
		middlewares.ErrorHandler(),
	)

	// group CRUD endpoints
	v1 := router.Group("/api/v1")
	{
		v1.GET("/err", func(c *gin.Context) {
			c.Error(fmt.Errorf("internal error"))
		})
	}

	router.GET("/api-health", handlers.SystemCheck)

	// Test endpoint to trigger error
	router.GET("/err", func(c *gin.Context) {
		c.Error(fmt.Errorf("Error Test"))
	})

	router.Run(port)
}
