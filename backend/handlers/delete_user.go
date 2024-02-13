package handlers

import (
	"fmt"
	"net/http"

	"github.com/1liale/maze-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RemoveUser(ctx *gin.Context) {
	dbCtx, exists := ctx.Get("db")
	if !exists {
		ctx.Error(fmt.Errorf("Error: DB object not found in ctx"))
		return
	}
	db, _ := dbCtx.(*gorm.DB)
	userID := ctx.Param("user")

	// get user corresponding to UserID
	var user models.User
	if err := db.Preload("Records").First(&user, "id = ?", userID).Error; err != nil {
		ctx.Error(err)
		return
	}
	if err := db.Delete(&models.MazeRecord{}, "user_id = ?", userID).Error; err != nil {
		ctx.Error(err)
		return
	}
	// Delete the user and its associated records
	if err := db.Delete(&user).Error; err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "SUCCESSFULLY REMOVED USER: " + userID})
}
