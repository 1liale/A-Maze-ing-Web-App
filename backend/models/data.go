package models

type MazeData struct {
	Start int   `json:"start" binding:"required"`
	End   int   `json:"end" binding:"required"`
	Grid  []int `json:"grid" binding:"required"`
}
type MazeSolution []int

type InputMazeBase struct {
	Width  int `json:"width" binding:"required,gte=3,lte=35"`
	Height int `json:"height" binding:"required,gte=3,lte=35"`
	Algorithms
}

type InputMazeSolve struct {
	InputMazeBase
	Maze     MazeData     `json:"maze" binding:"required"`
	Solution MazeSolution `json:"solution" binding:"required"`
}

type InputMazeSave struct {
	Maze      MazeData     `json:"maze" binding:"required"`
	Solution  MazeSolution `json:"solution" binding:"required"`
	SolveTime int          `json:"solve_time" binding:"required"`
}

type InputMazeDelete struct {
	MazeIDs []string `json:"maze_ids"`
}

type MazeOutputData struct {
	Maze    MazeData `json:"maze" binding:"required"`
	History [][2]int `json:"history" binding:"required"`
}

type MazeOutput struct {
	Data     MazeOutputData `json:"data" binding:"required"`
	Solution MazeSolution   `json:"solution" binding:"required"`
}

// extract output data
func ExtractMazeOutputData(m *Maze) MazeOutputData {
	output := MazeOutputData{
		Maze: MazeData{
			Start: m.Start.Pos,
			End:   m.End.Pos,
			Grid:  make([]int, m.N_Cells),
		},
		History: make([][2]int, len(m.History)),
	}
	for i, cell := range m.Cells {
		output.Maze.Grid[i] = cell.Wall
	}
	for i, cell := range m.History {
		output.History[i][0] = cell.Pos
		output.History[i][1] = cell.Wall
	}
	return output
}
