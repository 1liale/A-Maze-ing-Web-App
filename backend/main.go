// Package main is the root package managing the gin-gonic REST API.

// main.go
package main

import (
	"fmt"
	"os"

	"github.com/1liale/maze-backend/handlers"
	"github.com/1liale/maze-backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var logger *middlewares.CustomStdLogger
var db *gorm.DB

func init() {
	// load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}
	db = middlewares.InitDB()
	logger = middlewares.InitLogger()
}

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))

	// init gin server
	router := gin.New()

	// middlewares
	router.Use(
		ginlogrus.Logger(logger),
		gin.Recovery(),
		middlewares.ErrorHandler(),
		middlewares.PropDBEnv(db),
		cors.Default(), // enable CORS for all origins with all HTTP requests allowed by default
	)

	// require auth for requests that modify DB
	auth := router.Group("/").Use()
	{
		// create maze record for user (update if user exists already)
		auth.PUT("/maze/:user", handlers.SaveMaze)

		// delete user and their records
		auth.DELETE("/maze/:user", handlers.DeleteMaze)
	}

	// get records belonging to user
	router.GET("/maze/:user", handlers.GetMazes)

	// list number of maze records specified by the user or what's available, empty otherwise
	router.GET("/maze", handlers.ListMazes)

	// generate a new maze and corresponding solution with given user-specified dimensions
	router.POST("/maze/generate", handlers.GenerateMaze)

	// solve an unknown maze and send back solution
	router.POST("/maze/solve", handlers.SolveMaze)

	// gets a system check on api health
	auth.GET("/api-health", handlers.SystemCheck)

	router.Run(port)
}
