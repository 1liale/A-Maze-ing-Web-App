package services

import (
	"fmt"

	"github.com/1liale/maze-backend/models"
	"github.com/1liale/maze-backend/services/algos"
)

func MazeSolver(m *models.Maze, input *models.InputMazeBase, data *models.MazeData) (*models.Maze, error) {
	if m == nil {
		m = models.NewMaze(input.Width, input.Height)
		m.InitMaze(data)
	}

	// generate a maze using the specified algorithm
	solver := input.Algorithms.Solver
	fmt.Printf("Solving maze using algo: %s\n\n", solver)
	switch solver {
	case models.BFS:
		algos.BFS(m)
	case models.DFS:
		algos.DFS(m)
	case models.BBFS:
		algos.BBFS(m)
	}

	return m, nil
}
