package maze

type GenAlgorithm string

const (
	PrimAlgorithm    GenAlgorithm = "prim"
	KruskalAlgorithm GenAlgorithm = "kruskal"
)

type InputGenerateMaze struct {
	Width     int          `json:"width" binding:"required,gte=2,lte=10"`
	Height    int          `json:"height" binding:"required,gte=2,lte=10"`
	Algorithm GenAlgorithm `json:"algorithm" binding:"required,oneof=prim kruskal"`
}

type MazeOutput struct {
	Data     string
	Solution string
}
