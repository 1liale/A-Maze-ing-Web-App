package handlers

import (
	"fmt"

	"github.com/1liale/maze-backend/models/maze"
	"github.com/1liale/maze-backend/services"
	"github.com/gin-gonic/gin"
)

func GenerateMaze(ctx *gin.Context) {
	var maze_input maze.InputMazeBase
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

	m, err = services.MazeSolver(m, &maze_input, nil)
	if err != nil {
		ctx.Error(err)
		return
	}
	solution := maze.ExtractMazeOutputSoln(m)

	output := maze.MazeOutput{
		Data:     data,
		Solution: solution,
	}

	ctx.JSON(200, gin.H{"response": output})
}
