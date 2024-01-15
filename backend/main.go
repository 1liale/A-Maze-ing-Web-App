// Package main is the root package managing the gin-gonic REST API.

// main.go
package main

import (
	"net/http"

	"github.com/1liale/maze-backend/handlers"
	"github.com/1liale/maze-backend/middlewares"
	"github.com/1liale/maze-backend/models"
	ginlogrus "github.com/toorop/gin-logrus"

	"github.com/gin-gonic/gin"
)

var logger *middlewares.CustomStdLogger

func init() {
	logger = middlewares.NewCustomLogger()
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
		// search db for maze record matching username, empty otherwise
		v1.POST("/find-maze", handlers.FindMaze)

		// list number of maze records specified by the user or what's available, empty otherwise
		v1.POST("/list-mazes", handlers.ListMazes)

		// generate a new maze and corresponding solution with given user-specified dimensions
		v1.POST("/generate-maze", handlers.GenerateMaze)

		// solve an unknown maze and send back solution
		v1.POST("/solve-maze", handlers.SolveMaze)

		// save maze record (update if user exists already)
		v1.POST("/save-maze", handlers.SaveMaze)

		// delete maze record specified by the user, do nothing if not exist
		v1.DELETE("/delete-maze", handlers.DeleteMaze)
	}

	router.GET("/api-health", handlers.SystemCheck)

	// Test endpoint to trigger error
	router.GET("/err", func(c *gin.Context) {
		c.Error(&models.InternalError{
			Code: http.StatusBadGateway,
			Msg:  "Testing Error Endpoint",
		})
	})

	router.Run(port)
}
