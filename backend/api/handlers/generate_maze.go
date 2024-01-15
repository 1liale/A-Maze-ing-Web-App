package handlers

import (
	"fmt"

	"github.com/1liale/maze-backend/models/maze"
	"github.com/1liale/maze-backend/services"
	"github.com/gin-gonic/gin"
)

func GenerateMaze(ctx *gin.Context) {
	var maze_input maze.InputGenerateMaze
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
	data := maze.ExtractMazeOutputData(m)

	solution, err := services.MazeSolver(m, &maze_input, nil)

	output := maze.MazeOutput{
		Data:     data,
		Solution: solution,
	}

	ctx.JSON(200, gin.H{"data": output})
}
