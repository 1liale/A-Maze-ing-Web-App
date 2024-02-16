// Package main is the root package managing the gin-gonic REST API.

// main.go
package main

import (
	"context"
	"os"

	"github.com/1liale/maze-backend/config"
	"github.com/1liale/maze-backend/handlers"
	"github.com/1liale/maze-backend/middlewares"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	cors "github.com/rs/cors/wrapper/gin"

	ginlogrus "github.com/toorop/gin-logrus"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var logger *middlewares.CustomStdLogger
var db *gorm.DB
var conf *config.Config
var ginLambda *ginadapter.GinLambda

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Info("Error loading environment variables from .env, running ENV vars must be already set!")
	}

	env := os.Getenv("GIN_MODE")
	db = middlewares.InitDB()
	conf = config.InitConfig()
	router := gin.New()

	switch env {
	case "release":
		router.Use(
			ginlogrus.Logger(logrus.New()),
			gin.Recovery(),
			middlewares.ErrorHandler(),
			middlewares.PropDBEnv(db),
			cors.New(conf.CorsOptions),
		)
	default:
		logger = middlewares.InitLogger()
		// middlewares
		router.Use(
			ginlogrus.Logger(logger),
			gin.Recovery(),
			middlewares.ErrorHandler(),
			middlewares.PropDBEnv(db),
			// secure.Secure(conf.SecureOptions), TODO: enable this in production
			cors.New(conf.CorsOptions),
		)
	}

	router.GET("/api-health", handlers.SystemCheck)

	// require auth for requests that modify DB
	auth := router.Group("/").Use(middlewares.CheckJWT(conf.Audience, conf.Domain))
	{
		// create a new user if it doesn't exist in DB
		auth.GET("/user/:user", handlers.CheckCreateUser)
		// delete user from DB
		auth.DELETE("/user/:user", handlers.RemoveUser)
		// create maze record for user (update if user exists already)
		auth.PUT("/maze/:user", handlers.SaveMaze)
		// delete a user's maze records
		auth.DELETE("/maze/:user", handlers.DeleteMazes)
	}

	// get records belonging to user
	router.GET("/maze/:user", handlers.GetMazes)

	// list the top number of users by fastest solve time
	router.GET("/maze/scoreboard", handlers.GetScoreboard)

	// generate a new maze and corresponding solution with given user-specified dimensions
	router.POST("/maze/generate", handlers.GenerateMaze)

	// solve an unknown maze and send back solution
	router.POST("/maze/solve", handlers.SolveMaze)

	switch env {
	case "release":
		ginLambda = ginadapter.New(router)
	default:
		router.Run(conf.Port)
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	env := os.Getenv("GIN_MODE")
	if env == "release" {
		lambda.Start(Handler)
	}
}
