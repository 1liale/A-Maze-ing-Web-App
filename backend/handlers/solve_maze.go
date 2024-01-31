package handlers

import (
	"net/http"

	"github.com/1liale/maze-backend/models"
	"github.com/1liale/maze-backend/services"
	"github.com/gin-gonic/gin"
)

func SolveMaze(ctx *gin.Context) {
	var maze_input models.InputMazeSolve
	if err := ctx.ShouldBindJSON(&maze_input); err != nil {
		ctx.Error(err)
		return
	}

	m, err := services.MazeSolver(nil, &maze_input.InputMazeBase, &maze_input.Maze)
	if err != nil {
		ctx.Error(err)
		return
	}
	m.Display()

	data, solution := models.ExtractMazeOutputData(m), m.Path

	output := models.MazeOutput{
		Data:     data,
		Solution: solution,
	}

	ctx.JSON(http.StatusOK, gin.H{"response": output})
}
