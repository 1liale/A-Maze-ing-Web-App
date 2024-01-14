package services

import (
	"fmt"

	"github.com/1liale/maze-backend/models/maze"
	"github.com/1liale/maze-backend/services/algos"
)

func MazeGenerator(input *maze.InputGenerateMaze) (maze.MazeOutput, error) {
	// initialize a maze given input size
	m := maze.Maze{}
	m.InitMaze(input.Width, input.Height)

	// generate a maze using the specified algorithm
	fmt.Printf("Generating maze using algo: %s\n", input.Algorithm)
	switch input.Algorithm {
	case maze.PrimAlgorithm:
		algos.RandPrim(&m)
	case maze.KruskalAlgorithm:
		algos.RandKruskal(&m)
	}

	fmt.Printf("maze: %+v\n", m)

	return maze.MazeOutput{Data: "Hello", Solution: "World"}, nil
}
