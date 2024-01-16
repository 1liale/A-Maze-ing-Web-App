package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupDB(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	}
}
