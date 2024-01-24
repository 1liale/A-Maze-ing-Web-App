package algos

import (
	"github.com/1liale/maze-backend/models"
)

func addFrontiers(m *models.Maze, pos int, frontiers []int) []int {
	r, c := m.GetPosCoord(pos)
	// if m.Cells[pos].Wall&models.FRONTIER != 0 {
	// 	m.Cells[pos].Wall ^= models.FRONTIER
	// }
	m.Cells[pos].Wall |= models.IN
	for dir := 1; dir <= 8; dir *= 2 {
		nr, nc := r+models.Dr[dir], c+models.Dc[dir]

		if nr >= 0 && nr < m.Height && nc >= 0 && nc < m.Width {
			nCell := m.GetCellFromCoord(nr, nc)
			if nCell.Wall == 0 {
				nCell.Wall |= models.FRONTIER
				frontiers = append(frontiers, nCell.Pos)
			}
		}
	}
	return frontiers
}

func getNeighbours(m *models.Maze, pos int) []int {
	var neighbours []int
	r, c := m.GetPosCoord(pos)

	for dir := 1; dir <= 8; dir *= 2 {
		nr, nc := r+models.Dr[dir], c+models.Dc[dir]
		if nr >= 0 && nr < m.Height && nc >= 0 && nc < m.Width {
			if nCell := m.GetCellFromCoord(nr, nc); nCell.Wall&models.IN != 0 {
				neighbours = append(neighbours, nCell.Pos)
			}
		}
	}

	return neighbours
}

func getDirections(m *models.Maze, p1 int, p2 int) int {
	r, c := m.GetPosCoord(p1)
	nr, nc := m.GetPosCoord(p2)
	if r < nr {
		return models.D
	} else if r > nr {
		return models.U
	} else if c < nc {
		return models.R
	} else {
		return models.L
	}
}

func RandPrim(m *models.Maze) {
	pos := models.Rng.Intn(m.N_Cells)
	frontiers := addFrontiers(m, pos, []int{})

	for len(frontiers) > 0 {
		ind := models.Rng.Intn(len(frontiers))
		frontier := frontiers[ind]                                // get
		frontiers = append(frontiers[:ind], frontiers[ind+1:]...) // pop

		neighbours := getNeighbours(m, frontier)
		neighbour := neighbours[models.Rng.Intn(len(neighbours))]

		dir := getDirections(m, neighbour, frontier)
		// join walls between that frontier and the neighbour
		m.Cells[neighbour].Wall |= dir
		m.Cells[frontier].Wall |= models.Opp[dir]
		m.History = append(m.History, &m.Cells[neighbour])
		m.History = append(m.History, &m.Cells[frontier])

		frontiers = addFrontiers(m, frontier, frontiers)
	}
}
