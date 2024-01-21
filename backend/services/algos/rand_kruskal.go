package algos

import (
	"sort"

	"github.com/1liale/maze-backend/models"
)

func RandKruskal(m *models.Maze) {
	n_cells := m.Height * m.Width

	// build and sort edges by weight in ASC order
	edges := models.BuildEdges(m)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// initialize subsets, each cell has its own set initially
	subsets := make([]Subset, n_cells)
	for i := range subsets {
		subsets[i].Parent = i
		subsets[i].Rank = 0
	}

	newCells := append([]models.Cell{}, m.Cells...)
	m.History = append(m.History, newCells)
	for _, edge := range edges {
		src, dest := edge.Src, edge.Dest
		root1 := Find(subsets, src)
		root2 := Find(subsets, dest)
		dir := 0
		if src-dest == 1 {
			dir = models.L
		} else {
			dir = models.U
		}

		if root1 != root2 {
			Union(subsets, root1, root2)
			m.Cells[src].Wall |= dir
			m.Cells[dest].Wall |= models.Opp[dir]
			newCells = append([]models.Cell{}, m.Cells...)
			m.History = append(m.History, newCells)
		}
	}
}
