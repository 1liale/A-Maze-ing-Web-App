package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/1liale/maze-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListMazes(ctx *gin.Context) {
	dbCtx, exists := ctx.Get("db")
	if !exists {
		ctx.Error(fmt.Errorf("Error: DB object not found in ctx"))
		return
	}
	db, _ := dbCtx.(*gorm.DB)

	limit_str := ctx.DefaultQuery("limit", "5")
	limit, err := strconv.Atoi(limit_str)
	if err != nil || limit < 0 {
		ctx.Error(fmt.Errorf("Query string 'limit' must be an integer > 0"))
		return
	}

	// Retrieve the top 5 records sorted by SolveTime
	var maze_records []models.MazeRecord
	db.Order("solve_time").Limit(limit).Find(&maze_records)

	ctx.JSON(http.StatusOK, gin.H{"response": maze_records})
}
