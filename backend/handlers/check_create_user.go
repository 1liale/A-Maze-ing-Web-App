package handlers

import (
	"fmt"
	"net/http"

	"github.com/1liale/maze-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckCreateUser(ctx *gin.Context) {
	dbCtx, exists := ctx.Get("db")
	if !exists {
		ctx.Error(fmt.Errorf("Error: DB object not found in ctx"))
		return
	}
	db, _ := dbCtx.(*gorm.DB)

	userID := ctx.Param("user")
	userName := ctx.Query("name")
	if len(userName) == 0 {
		ctx.Error(fmt.Errorf("Error: User has to have a non-empty name"))
		return
	}

	// get user corresponding to UserID
	var user models.User
	result := db.First(&user, "id = ?", userID)
	if result.Error != nil {
		// create user
		user = models.User{ID: userID, Name: userName, Highscore: 0}
		db.Create(&user)
		ctx.JSON(http.StatusOK, gin.H{"response": "SUCCESSFULLY CREATED USER: " + userID})
	} else {
		if userName != user.Name {
			user.Name = userName
			db.Save(&user)
		}
		ctx.JSON(http.StatusOK, gin.H{"response": "USER EXISTS!"})
	}

}
