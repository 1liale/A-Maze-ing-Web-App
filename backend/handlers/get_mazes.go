package handlers

import (
	"fmt"
	"net/http"

	"github.com/1liale/maze-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMazes(ctx *gin.Context) {
	dbCtx, exists := ctx.Get("db")
	if !exists {
		ctx.Error(fmt.Errorf("Error: DB object not found in ctx"))
		return
	}
	db, _ := dbCtx.(*gorm.DB)

	userID := ctx.Param("user")

	// Retrieve the top 5 records sorted by SolveTime
	var maze_records []models.MazeRecord
	db.Preload("Records").Where("user_id = ?", userID).Order("score").Find(&maze_records)

	ctx.JSON(http.StatusOK, gin.H{"response": maze_records})
}
