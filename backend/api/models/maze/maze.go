package maze

import (
	"math/rand"
)

// rand
const SEED int64 = 42

var rnd *rand.Rand = rand.New(rand.NewSource(SEED))

type Dir int

// setup constants for wall directions
const (
	U, D, L, R = 1, 2, 4, 8
	CLOSED     = 15 // assume all walls are closed initially
)

// directional mapping for easier calc
var (
	Dr  = map[Dir]int{L: 0, R: 0, U: -1, D: 1}
	Dc  = map[Dir]int{L: -1, R: 1, U: 0, D: 0}
	Opp = map[Dir]Dir{L: R, R: L, U: D, D: U}
)

// define maze data structures
type Cell struct {
	Pos  int
	Wall Dir
}

// common maze struct for all algorithms
type Maze struct {
	Width   int
	Height  int
	N_Cells int
	Cells   []Cell
}

// initialize maze
func (m *Maze) InitMaze(w, h int) {
	m.N_Cells = w * h
	cells := make([]Cell, m.N_Cells)
	for i := 0; i < m.N_Cells; i++ {
		cells[i].Pos = i
		cells[i].Wall = 0
	}

	m.Width, m.Height = w, h
	m.Cells = cells
}

// get a cell's Coordinate
func (m *Maze) GetPosCoord(pos int) (int, int) {
	return pos / m.Width, pos % m.Width
}

type Edge struct {
	Src, Dest int
	Weight    int
}

type Graph struct {
	Edges    []Edge
	Vertices int
}

// build edges given a maze / matrix
func BuildEdges(m *Maze) []Edge {
	w, h := m.Width, m.Height
	n_edges := 2*w*h - w - h // number of edges in a 2d lattice
	edges := make([]Edge, n_edges)

	edge_ind := 0
	for i, cell := range m.Cells {
		r, c := m.GetPosCoord(cell.Pos)
		if r > 0 {
			edges[edge_ind] = Edge{
				Src:    i,
				Dest:   (r+1)*w + c,
				Weight: rnd.Intn(h * w),
			}
			edge_ind++
		}
		if c > 0 {
			edges[edge_ind] = Edge{
				Src:    i,
				Dest:   r*w + c + 1,
				Weight: rnd.Intn(h * w),
			}
			edge_ind++
		}
	}
	return edges
}
