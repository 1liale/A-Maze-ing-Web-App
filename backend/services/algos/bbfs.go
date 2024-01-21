package algos

import (
	"fmt"
	"time"

	"github.com/1liale/maze-backend/models"
)

// pathFrom reads the provided predecessor map to build a path from u to v. If pred contains
// a cycle, this method will detect it and return an error.
func pathFrom(pred map[int]int, u int, v int) []int {
	var result []int
	seen := make(map[int]bool, len(pred))
	curr := u
	for {
		if seen[curr] {
			// return with error
			return nil
		}
		result = append(result, curr)
		if curr == v {
			break
		}
		seen[curr] = true
		curr = pred[curr]
	}
	return result
}

func goBFS(start int, done chan int, peerOut chan<- int, peerIn <-chan int, m *models.Maze) <-chan []int {
	result := make(chan []int, 1)
	go func() {
		frontier := []int{start}      // my search frontier
		visited := make(map[int]bool) // my visited array
		pred := make(map[int]int)     // predecessors I know about
		edgeCount := 0
		defer close(result)
		defer close(peerOut)
	finish:
		for {
			select {
			case meetNode := <-done:
				// peer found the meeting point, send back our half of the path; note pred is guaranteed
				// to contain a valid path, since peer found a meeting point based on a node we visited.
				result <- pathFrom(pred, meetNode, start)
				break finish
			case other := <-peerIn: // peer has visited another node
				if visited[other] {
					// search frontiers have merged, send back our half of the path.
					done <- other
					result <- pathFrom(pred, other, start)
					break finish
				}
			default:
				// expand search and send peer the visited node
				next := frontier[0]
				frontier = frontier[1:]
				if !visited[next] {
					// neighbors := graph[next]
					var neighbors []int
					for i := 0; i < len(neighbors); i++ {
						if _, iVisited := visited[neighbors[i]]; !iVisited {
							edgeCount++
							pred[neighbors[i]] = next
							frontier = append(frontier, neighbors[i])
							peerOut <- neighbors[i]
						}
					}
					visited[next] = true
				}

				if len(frontier) == 0 {
					result <- nil // no path exists; report this fact
					break finish  // terminate search once the frontier is empty
				}
			}
		}
		fmt.Println("BFS starting from", start, "visited", len(visited), "nodes and", edgeCount, "edges")
	}()
	return result
}

func BBFS(m *models.Maze) []int {
	done := make(chan int, 1)
	uChan := make(chan int, 1000)
	vChan := make(chan int, 1000)

	start := time.Now()
	uResult := goBFS(m.Start.Pos, done, vChan, uChan, m)
	vResult := goBFS(m.End.Pos, done, uChan, vChan, m)

	var uPath, vPath []int
	select {
	case uPath = <-uResult:
		if uPath == nil {
			return nil
		}
		vPath = <-vResult
	case vPath = <-vResult:
		if vPath == nil {
			return nil
		}
		uPath = <-uResult
	}

	result := append(reverse(uPath), vPath[1:]...)

	fmt.Println("Search took", time.Since(start))

	return result
}

func reverse(slice []int) []int {
	n := len(slice)
	for i := 0; i < n/2; i++ {
		slice[i], slice[n-i-1] = slice[n-i-1], slice[i]
	}
	return slice
}
