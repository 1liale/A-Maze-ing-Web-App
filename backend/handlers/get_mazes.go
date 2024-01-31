package handlers

import (
	"fmt"
	"net/http"

	"github.com/1liale/maze-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetMazes(ctx *gin.Context) {
	dbCtx, exists := ctx.Get("db")
	if !exists {
		ctx.Error(fmt.Errorf("Error: DB object not found in ctx"))
		return
	}
	db, _ := dbCtx.(*gorm.DB)

	userIDStr := ctx.Param("user")
	userID, err := uuid.Parse(userIDStr)
	// should never have an invalid uuid for an auth user
	if err != nil {
		ctx.Error(fmt.Errorf("Invalid user ID"))
		return
	}

	// Retrieve the top 5 records sorted by SolveTime
	var maze_records []models.MazeRecord
	db.Where("user_id = ?", userID).Order("solve_time").Find(&maze_records)

	ctx.JSON(http.StatusOK, gin.H{"response": maze_records})
}
