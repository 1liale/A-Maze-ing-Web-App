// Package main is the root package managing the gin-gonic REST API.

// main.go
package main

import (
	"github.com/1liale/maze-backend/config"
	"github.com/1liale/maze-backend/handlers"
	"github.com/1liale/maze-backend/middlewares"

	cors "github.com/rs/cors/wrapper/gin"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var logger *middlewares.CustomStdLogger
var db *gorm.DB
var conf *config.Config

func init() {
	// load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}
	db = middlewares.InitDB()
	logger = middlewares.InitLogger()
	conf = config.InitConfig()
}

func main() {
	// init gin server
	router := gin.New()

	// middlewares
	router.Use(
		ginlogrus.Logger(logger),
		gin.Recovery(),
		middlewares.ErrorHandler(),
		middlewares.PropDBEnv(db),
		// secure.Secure(conf.SecureOptions), TODO: enable this in production
		cors.New(conf.CorsOptions),
	)

	// require auth for requests that modify DB
	auth := router.Group("/").Use(middlewares.CheckJWT(conf.Audience, conf.Domain))
	{
		// create maze record for user (update if user exists already)
		auth.PUT("/maze/:user", handlers.SaveMaze)
		// delete a user's maze records
		auth.DELETE("/maze/:user", handlers.DeleteMazes)

		auth.GET("/api-health", handlers.SystemCheck)
	}

	// get records belonging to user
	router.GET("/maze/:user", handlers.GetMazes)

	// list number of maze records specified by the user or what's available, empty otherwise
	router.GET("/maze", handlers.ListMazes)

	// generate a new maze and corresponding solution with given user-specified dimensions
	router.POST("/maze/generate", handlers.GenerateMaze)

	// solve an unknown maze and send back solution
	router.POST("/maze/solve", handlers.SolveMaze)

	router.Run(conf.Port)
}
