package maze

type Algorithm string

// Generator Algorithms
const (
	Prim    Algorithm = "prim"
	Kruskal Algorithm = "kruskal"
)

// Solver Algorithms
const (
	BFS  Algorithm = "bfs"
	BBFS Algorithm = "bbfs"
	DFS  Algorithm = "dfs"
)

type Algorithms struct {
	Generator Algorithm `json:"generator" binding:"required,oneof=prim kruskal"`
	Solver    Algorithm `json:"solver" binding:"required,oneof=bfs dfs bbfs"`
}
