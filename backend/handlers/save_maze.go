package handlers

import (
	"fmt"
	"net/http"

	"github.com/1liale/maze-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveMaze(ctx *gin.Context) {
	dbCtx, exists := ctx.Get("db")
	if !exists {
		ctx.Error(fmt.Errorf("Error: DB object not found in ctx"))
		return
	}
	db, _ := dbCtx.(*gorm.DB)

	userID := ctx.Param("user")

	// get user corresponding to UserID
	var user models.User
	result := db.First(&user, "id = ?", userID)
	if result.Error != nil {
		ctx.Error(fmt.Errorf("No USER with that ID exists"))
		return
	}

	// parse JSON body
	var maze_input models.InputMazeSave
	if err := ctx.ShouldBindJSON(&maze_input); err != nil {
		ctx.Error(err)
		return
	}

	// update user fields
	if user.Highscore == 0 || maze_input.Score > user.Highscore {
		user.Highscore = maze_input.Score
	}
	db.Save(&user)

	// create new maze record associated with user
	maze_record := models.MazeRecord{
		UserId:   userID,
		UserName: user.Name,
		Data:     maze_input.Maze,
		Solution: maze_input.Solution,
		Score:    maze_input.Score,
	}
	db.Create(&maze_record)

	ctx.JSON(http.StatusOK, gin.H{"response": "SUCCESSFULLY UPDATED MAZE RECORDS FOR USER: " + userID})
}
