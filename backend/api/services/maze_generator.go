package services

import (
	"fmt"

	"github.com/1liale/maze-backend/models/maze"
	"github.com/1liale/maze-backend/services/algos"
)

func MazeGenerator(input *maze.InputMazeBase) (*maze.Maze, error) {
	// initialize a maze given input size
	m := maze.NewMaze(input.Width, input.Height)
	m.SetRandomEndpoints()

	// generate a maze using the specified algorithm
	generator := input.Algorithms.Generator
	fmt.Printf("Generating maze using algo: %s\n\n", generator)
	switch generator {
	case maze.Prim:
		algos.RandPrim(m)
	case maze.Kruskal:
		algos.RandKruskal(m)
	}

	// Display "perfect" maze
	m.Display()
	return m, nil
}
