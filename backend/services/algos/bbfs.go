package algos

import (
	"slices"
	"sync"

	"github.com/1liale/maze-backend/models"
)

func BBFS(m *models.Maze) {
	done := make(chan int, 1)
	flag1 := make(chan int, m.N_Cells)
	flag2 := make(chan int, m.N_Cells)

	paths := make(chan []int, 2) // Creating a buffered channel
	var wg sync.WaitGroup

	BFSHelper(m, m.Start.Pos, &wg, paths, done, flag2, flag1)
	BFSHelper(m, m.End.Pos, &wg, paths, done, flag1, flag2)

	go func() {
		wg.Wait()
		close(paths) // Close the channel when all paths have finished
	}()

	pathMap := make(map[int]models.MazeSolution)
	for path := range paths {
		if path != nil {
			pathMap[path[0]] = path
		}
	}

	startPath := pathMap[m.Start.Pos]
	endPath, exists := pathMap[m.End.Pos]
	if exists && len(endPath) > 1 {
		slices.Reverse(endPath)
		m.Path = append(startPath, endPath[1:]...)
	} else {
		m.Path = startPath
	}
}
