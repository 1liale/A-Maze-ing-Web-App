package handlers

import (
	"fmt"

	"github.com/1liale/maze-backend/models"
	"github.com/1liale/maze-backend/services"
	"github.com/gin-gonic/gin"
)

func GenerateMaze(ctx *gin.Context) {
	var maze_input models.InputMazeBase
	if err := ctx.ShouldBindJSON(&maze_input); err != nil {
		ctx.Error(err)
		return
	}

	fmt.Printf("Received Request Body: %+v\n", maze_input)
	m, err := services.MazeGenerator(&maze_input)
	if err != nil {
		ctx.Error(err)
		return
	}
	data := models.ExtractMazeOutputData(m)

	m, err = services.MazeSolver(m, &maze_input, nil)
	if err != nil {
		ctx.Error(err)
		return
	}
	solution := models.ExtractMazeOutputSoln(m)

	output := models.MazeOutput{
		Data:     data,
		Solution: solution,
	}

	ctx.JSON(200, gin.H{"response": output})
}
