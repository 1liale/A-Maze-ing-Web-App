package algos

import (
	"sort"

	"github.com/1liale/maze-backend/models/maze"
)

func RandKruskal(m *maze.Maze) {
	n_cells := m.Height * m.Width

	// build and sort edges by weight in ASC order
	edges := maze.BuildEdges(m)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// initialize subsets, each cell has its own set initially
	subsets := make([]Subset, n_cells)
	for i := range subsets {
		subsets[i].Parent = i
		subsets[i].Rank = 0
	}

	for _, edge := range edges {
		src, dest := edge.Src, edge.Dest
		root1 := Find(subsets, src)
		root2 := Find(subsets, dest)
		var dir maze.Dir
		if src/m.Width == dest/m.Width {
			dir = maze.L
		} else {
			dir = maze.U
		}

		if root1 != root2 {
			Union(subsets, root1, root2)
			m.Cells[src].Wall |= dir
			m.Cells[dest].Wall |= maze.Opp[dir]
		}
	}
}
