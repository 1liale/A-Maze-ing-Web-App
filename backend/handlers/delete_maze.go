package handlers

import (
	"fmt"
	"net/http"

	"github.com/1liale/maze-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteMazes(ctx *gin.Context) {
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

	// parse JSON body
	var maze_input models.InputMazeDelete
	if err := ctx.ShouldBindJSON(&maze_input); err != nil {
		ctx.Error(err)
		return
	}
	maze_ids := maze_input.MazeIDs

	var result *gorm.DB
	if len(maze_ids) > 0 {
		result = db.Where("user_id = ? AND id IN ?", userID, maze_ids).Delete(&models.MazeRecord{})
	} else {
		result = db.Where("user_id = ?", userID).Delete(&models.MazeRecord{})
	}

	if result.Error != nil {
		ctx.Error(fmt.Errorf("Error deleting records: %v", result.Error))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "SUCCESSFULLY DELETED MAZE RECORDS FOR USER: " + userIDStr})
}
