package services

import (
	"fmt"

	"github.com/1liale/maze-backend/models"
	"github.com/1liale/maze-backend/services/algos"
)

func MazeGenerator(input *models.InputMazeBase) (*models.Maze, error) {
	// initialize a maze given input size
	m := models.NewMaze(input.Width, input.Height)
	m.SetRandomEndpoints()

	// generate a maze using the specified algorithm
	generator := input.Algorithms.Generator
	fmt.Printf("Generating maze using algo: %s\n\n", generator)
	switch generator {
	case models.Prim:
		algos.RandPrim(m)
	case models.Kruskal:
		algos.RandKruskal(m)
	}

	// Display "perfect" maze
	m.Display()
	return m, nil
}
