// Package main is the root package managing the gin-gonic REST API.

// main.go
package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"

	"github.com/1liale/maze-backend/handlers"
	"github.com/1liale/maze-backend/middlewares"
	"github.com/1liale/maze-backend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

type Product struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Code  string  `json:"code"`
	Price float32 `json:"price"`
}

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))

	// init gin server
	router := gin.New()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))

	// middlewares
	router.Use(
		ginlogrus.Logger(logger),
		gin.Recovery(),
		middlewares.ErrorHandler(),
		middlewares.PropDBEnv(db),
		cors.Default(), // enable CORS for all origins with all HTTP requests allowed by default
		sessions.Sessions("maze-session", store),
	)

	authenticator, err := models.NewAuth()
	if err != nil {
		logger.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	// TODO: Remove when done testing
	router.GET("/", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
		<html><div>
			<h3>Auth0 Example</h3>
			<a href="/login">SignIn</a>
		</div></html>`))
	})

	// special routes (auth, system checks, etc.)
	router.GET("/login", handlers.Login(authenticator))
	router.GET("/logout", handlers.Logout)
	router.GET("/callback", handlers.Callback(authenticator))
	router.GET("/api-health", handlers.SystemCheck)

	// require auth for requests that modify DB
	auth := router.Group("/").Use(middlewares.IsAuthenticated)
	{
		// create maze record for user (update if user exists already)
		auth.PUT("/maze/:user", handlers.SaveMaze)
		// delete a user's maze records
		auth.DELETE("/maze/:user", handlers.DeleteMazes)
	}

	// get records belonging to user
	router.GET("/maze/:user", handlers.GetMazes)

	// list number of maze records specified by the user or what's available, empty otherwise
	router.GET("/maze", handlers.ListMazes)

	// generate a new maze and corresponding solution with given user-specified dimensions
	router.POST("/maze/generate", handlers.GenerateMaze)

	// solve an unknown maze and send back solution
	router.POST("/maze/solve", handlers.SolveMaze)

	router.Run(port)
}
