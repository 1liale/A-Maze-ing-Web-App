package handlers

import (
	"github.com/1liale/maze-backend/models/maze"
	"github.com/1liale/maze-backend/services"
	"github.com/gin-gonic/gin"
)

func SolveMaze(ctx *gin.Context) {
	var maze_input maze.InputMazeSolve
	if err := ctx.ShouldBindJSON(&maze_input); err != nil {
		ctx.Error(err)
		return
	}

	m, err := services.MazeSolver(nil, &maze_input.InputMazeBase, &maze_input.Maze)
	if err != nil {
		ctx.Error(err)
		return
	}

	data, solution := maze.ExtractMazeOutputData(m), maze.ExtractMazeOutputSoln(m)

	output := maze.MazeOutput{
		Data:     data,
		Solution: solution,
	}

	ctx.JSON(200, gin.H{"response": output})
}
