package algos

import (
	"fmt"

	"github.com/1liale/maze-backend/models"
	"github.com/adrianbrad/queue"
)

func rebuildPath(m *models.Maze, parents map[int]int, pos int) {
	var path []*models.Cell

	for pos != m.Start.Pos {
		if pos != m.End.Pos {
			m.Cells[pos].State = "*"
		}
		path = append([]*models.Cell{&m.Cells[pos]}, path...)
		pos = parents[pos]
	}
	path = append([]*models.Cell{&m.Cells[pos]}, path...)
	m.Path = path

	fmt.Print("path:")
	for _, p := range path {
		fmt.Printf(" %d", p.Pos)
	}
	fmt.Print("\n\n")
	m.Display()
}

func BFS(m *models.Maze) {
	// queue of cell positions
	q := queue.NewLinked([]int{})
	visited := make([]bool, m.N_Cells)
	parents := make(map[int]int)

	q.Offer(m.Start.Pos)

	for !q.IsEmpty() {
		pos, _ := q.Get()
		// found the maze's exit
		if pos == m.End.Pos {
			// reconstruct path
			rebuildPath(m, parents, pos)
			return
		}

		visited[pos] = true
		for _, neighbour := range m.GetNeighbours(pos) {
			if !visited[neighbour.Pos] {
				parents[neighbour.Pos] = pos
				visited[neighbour.Pos] = true
				q.Offer(neighbour.Pos)
			}
		}
	}
}
