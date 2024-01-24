package algos

import "github.com/1liale/maze-backend/models"

func RebuildPath(parents []int, pos int, start int) models.MazeSolution {
	var path models.MazeSolution

	for pos != start {
		path = append(models.MazeSolution{pos}, path...)
		pos = parents[pos]
	}
	path = append(models.MazeSolution{start}, path...)

	return path
}
