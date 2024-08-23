package middlewares

import (
	"os"

	"github.com/1liale/maze-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PropDBEnv(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
		
	}
}

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Warn("Cannot connect to: ", dsn)
		logrus.Info("Proceeding without db connection!")
	} else {
		logrus.Info("Successfully connected to: ", dsn)
		db.AutoMigrate(&models.User{}, &models.MazeRecord{})
	}

	return db
}
