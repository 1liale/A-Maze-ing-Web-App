package maze

type InputGenerateMaze struct {
	Width      int `json:"width" binding:"required,gte=2,lte=10"`
	Height     int `json:"height" binding:"required,gte=2,lte=10"`
	Algorithms Algorithms
}

type MazeData struct {
	Start int
	End   int
	Grid  []int // binary data about which walls are open
}

type MazeOutputData struct {
	Maze    MazeData
	History [][]int
}

// extract output data
func ExtractMazeOutputData(m *Maze) MazeOutputData {
	output := MazeOutputData{
		Maze: MazeData{
			Start: m.Start.Pos,
			End:   m.End.Pos,
			Grid:  make([]int, m.N_Cells),
		},
		History: make([][]int, len(m.History)),
	}
	for i, cell := range m.Cells {
		output.Maze.Grid[i] = cell.Wall
	}
	for i := range m.History {
		tmp := make([]int, m.N_Cells)
		for j, cell := range m.History[i] {
			tmp[j] = cell.Wall
		}
		output.History[i] = tmp
	}
	return output
}

type MazeOutputSolution []int

func ExtractMazeOutputSoln(m *Maze) MazeOutputSolution {
	solution := make(MazeOutputSolution, len(m.Path))
	for i, cell := range m.Path {
		solution[i] = cell.Pos
	}
	return solution
}

type MazeOutput struct {
	Data     MazeOutputData
	Solution MazeOutputSolution
}
