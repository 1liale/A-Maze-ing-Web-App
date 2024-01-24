package algos

import (
	"github.com/1liale/maze-backend/models"
)

func dfsHelper(m *models.Maze, visited []bool, pos int, stack []int) {
	stack = append(stack, pos)

	visited[pos] = true
	if pos == m.End.Pos {
		m.Path = stack
	}

	for _, neighbour := range m.GetNeighbours(pos) {
		if !visited[neighbour.Pos] {
			dfsHelper(m, visited, neighbour.Pos, append([]int{}, stack...))
		}
	}
}

func DFS(m *models.Maze) {
	visited := make([]bool, m.N_Cells)
	start := m.Start.Pos
	stack := make([]int, 0)
	dfsHelper(m, visited, start, stack)
}
