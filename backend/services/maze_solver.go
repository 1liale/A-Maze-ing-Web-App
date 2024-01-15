package services

import (
	"fmt"

	"github.com/1liale/maze-backend/models/maze"
	"github.com/1liale/maze-backend/services/algos"
)

func MazeSolver(m *maze.Maze, input *maze.InputMazeBase, data *maze.MazeData) (*maze.Maze, error) {
	if m == nil {
		m = maze.NewMaze(input.Width, input.Height)
		m.InitMaze(data)
	}

	// generate a maze using the specified algorithm
	solver := input.Algorithms.Solver
	fmt.Printf("Solving maze using algo: %s\n\n", solver)
	switch solver {
	case maze.BFS:
		algos.BFS(m)
	case maze.DFS:
		algos.DFS(m)
	case maze.BBFS:
		algos.BBFS(m)
	}

	return m, nil
}
