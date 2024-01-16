// Package main is the root package managing the gin-gonic REST API.

// main.go
package main

import (
	"github.com/1liale/maze-backend/handlers"
	"github.com/1liale/maze-backend/middlewares"
	"github.com/1liale/maze-backend/models"
	"github.com/gin-contrib/cors"
	"github.com/jinzhu/gorm"
	ginlogrus "github.com/toorop/gin-logrus"

	"github.com/gin-gonic/gin"
)

var logger *middlewares.CustomStdLogger
var db *gorm.DB

func init() {
	logger = middlewares.NewCustomLogger()
	db = models.SetupModels()
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
		cors.Default(), // enable CORS for all origins with all HTTP requests allowed by default
		middlewares.SetupDB(db),
	)

	// require auth for requests that modify DB
	v1 := router.Group("/auth")
	{
		// save maze record (update if user exists already)
		v1.POST("/save-maze", handlers.SaveMaze)

		// delete maze record specified by the user, do nothing if not exist
		v1.DELETE("/delete-maze", handlers.DeleteMaze)
	}

	// search db for maze record matching username, empty otherwise
	router.POST("/find-maze", handlers.FindMaze)

	// list number of maze records specified by the user or what's available, empty otherwise
	router.POST("/list-mazes", handlers.ListMazes)

	// generate a new maze and corresponding solution with given user-specified dimensions
	router.POST("/generate-maze", handlers.GenerateMaze)

	// solve an unknown maze and send back solution
	router.POST("/solve-maze", handlers.SolveMaze)

	// gets a system check on api health
	router.GET("/api-health", handlers.SystemCheck)

	// TODO: Remove, for testing purpose only
	router.GET("/books", handlers.FindBooks)
	router.POST("/books", handlers.CreateBook)       // create
	router.GET("/books/:id", handlers.FindBook)      // find by id
	router.PATCH("/books/:id", handlers.UpdateBook)  // update by id
	router.DELETE("/books/:id", handlers.DeleteBook) // delete by id

	router.Run(port)
}
