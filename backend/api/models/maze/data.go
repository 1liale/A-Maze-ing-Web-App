package maze

type MazeData struct {
	Start int   `json:"start" binding:"required"`
	End   int   `json:"end" binding:"required"`
	Grid  []int `json:"grid" binding:"required"`
}
type MazeSolution []int

type InputMazeBase struct {
	Width  int `json:"width" binding:"required,gte=2,lte=10"`
	Height int `json:"height" binding:"required,gte=2,lte=10"`
	Algorithms
}

type InputMazeSolve struct {
	InputMazeBase
	Maze     MazeData     `json:"maze" binding:"required"`
	Solution MazeSolution `json:"solution" binding:"required"`
}

type MazeOutputData struct {
	Maze    MazeData `json:"maze" binding:"required"`
	History [][]int  `json:"history" binding:"required"`
}

type MazeOutputSolution MazeSolution
type MazeOutput struct {
	Data     MazeOutputData     `json:"data" binding:"required"`
	Solution MazeOutputSolution `json:"solution" binding:"required"`
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

func ExtractMazeOutputSoln(m *Maze) MazeOutputSolution {
	solution := make(MazeOutputSolution, len(m.Path))
	for i, cell := range m.Path {
		solution[i] = cell.Pos
	}
	return solution
}
