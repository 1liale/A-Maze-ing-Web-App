package models

import (
	"fmt"
	"math/rand"
)

// rand
const SEED int64 = 42

var rnd *rand.Rand = rand.New(rand.NewSource(SEED))

// setup constants for wall directions
const (
	U, D, L, R = 1, 2, 4, 8
	CLOSED     = 15 // assume all walls are closed initially
)

// directional mapping for easier calc
var (
	Dr  = map[int]int{U: -1, D: 1, L: 0, R: 0}
	Dc  = map[int]int{U: 0, D: 0, L: -1, R: 1}
	Opp = map[int]int{L: R, R: L, U: D, D: U}
)

type Edge struct {
	Src, Dest int
	Weight    int
}

type Graph struct {
	Edges    []Edge
	Vertices int
}

// define maze data structures
type Cell struct {
	Pos   int
	Wall  int
	State string
}

func NewCell(pos int) Cell {
	return Cell{Pos: pos, Wall: 0, State: "."}
}

// common maze struct for all algorithms
type Maze struct {
	Width   int
	Height  int
	N_Cells int
	Start   *Cell
	End     *Cell
	Path    []*Cell
	Cells   []Cell
	History [][]Cell // to show generator progress over iterations, can visualize
}

// pick start randomly on left wall and end randomly from right wall
func (m *Maze) SetRandomEndpoints() {
	w, h := m.Width, m.Height
	r1 := rnd.Intn(h)
	r2 := rnd.Intn(h)

	m.End = m.GetCellFromCoord(r2, w-1)
	m.Start = m.GetCellFromCoord(r1, 0)

	m.Start.State = "S"
	m.End.State = "E"
}

type MazeConstructor interface{}

// maze constructor
func NewMaze(w int, h int) *Maze {
	m := Maze{}
	m.Width, m.Height = w, h
	m.N_Cells = w * h
	cells := make([]Cell, m.N_Cells)
	for i := 0; i < m.N_Cells; i++ {
		cells[i] = NewCell(i)
	}
	m.Cells = cells
	return &m
}

// populate with maze data
func (m *Maze) InitMaze(data *MazeData) {
	for i := 0; i < m.N_Cells; i++ {
		m.Cells[i].Wall = data.Grid[i]
	}
	m.Start = &m.Cells[data.Start]
	m.End = &m.Cells[data.End]

	m.Start.State = "S"
	m.End.State = "E"
}

// get a cell's Coordinate
func (m *Maze) GetPosCoord(pos int) (int, int) {
	return pos / m.Width, pos % m.Width
}

// get a cell given its Coordinate
func (m *Maze) GetCellFromCoord(r, c int) *Cell {
	return &m.Cells[r*m.Width+c]
}

// get neighbours
func (m *Maze) GetNeighbours(pos int) []Cell {
	r, c := m.GetPosCoord(pos)
	cell := m.Cells[pos]
	var neighbours []Cell

	for dir := 1; dir <= 8; dir *= 2 {
		nr, nc := r+Dr[dir], c+Dc[dir]

		if nr >= 0 && nr < m.Height && nc >= 0 && nc < m.Width && dir&cell.Wall != 0 {
			neighbours = append(neighbours, *m.GetCellFromCoord(nr, nc))
		}
	}

	return neighbours
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
				Dest:   (r-1)*w + c,
				Weight: rnd.Intn(h * w),
			}
			edge_ind++
		}
		if c > 0 {
			edges[edge_ind] = Edge{
				Src:    i,
				Dest:   r*w + c - 1,
				Weight: rnd.Intn(h * w),
			}
			edge_ind++
		}
	}
	return edges
}

// display the maze inside the console
func (m *Maze) Display() {
	for c := 0; c < m.Width; c++ {
		fmt.Print("+---")
	}
	fmt.Println("+")
	for r := 0; r < m.Height; r++ {
		fmt.Print("+")
		for c := 0; c < m.Width; c++ {
			cell := m.GetCellFromCoord(r, c)
			if cell.Wall&R == 0 {
				fmt.Printf(" %v +", cell.State)
			} else {
				fmt.Printf(" %v  ", cell.State)
			}
		}
		fmt.Println()
		fmt.Print("+")
		for c := 0; c < m.Width; c++ {
			cell := m.GetCellFromCoord(r, c)
			cellStr := ""
			if cell.Wall&D == 0 {
				cellStr += "---+"
			} else {
				cellStr += "   +"
			}
			fmt.Print(cellStr)
		}
		fmt.Println()
	}
	fmt.Println()
}
