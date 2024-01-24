package algos

import (
	"sync"
	"time"

	"github.com/1liale/maze-backend/models"
	"github.com/adrianbrad/queue"
)

func BFSHelper(m *models.Maze, start int, wg *sync.WaitGroup, paths chan []int, flags ...chan int) {
	var doneFlag, peerOut, peerIn chan int

	argsLen := len(flags)
	if argsLen != 0 && argsLen != 3 {
		return
	} else if argsLen == 3 {
		doneFlag = flags[0]
		peerOut = flags[1]
		peerIn = flags[2]
	}

	wg.Add(1)
	go func() {
		q := queue.NewLinked([]int{start})
		visited := make([]bool, m.N_Cells)
		parents := make([]int, m.N_Cells)
		defer wg.Done()
	done:
		for {
			time.Sleep(20 * time.Microsecond)
			select {
			case peerIntersectNode := <-doneFlag:
				paths <- RebuildPath(parents, peerIntersectNode, start)
				break done
			case peerCurrentNode := <-peerIn:
				if visited[peerCurrentNode] {
					doneFlag <- peerCurrentNode
					paths <- RebuildPath(parents, peerCurrentNode, start)
					break done
				}
			default:
				currentNode, _ := q.Get()
				// solution found by single thread BFS
				if argsLen == 0 && (currentNode == m.Start.Pos || currentNode == m.End.Pos) && currentNode != start {
					paths <- RebuildPath(parents, currentNode, start)
					break done
				}

				if !visited[currentNode] {
					for _, neighbour := range m.GetNeighbours(currentNode) {
						if !visited[neighbour.Pos] {
							parents[neighbour.Pos] = currentNode
							q.Offer(neighbour.Pos)
							if argsLen == 3 {
								peerOut <- neighbour.Pos
							}
						}
					}
					visited[currentNode] = true
				}

				if q.IsEmpty() {
					paths <- nil
					break done
				}
			}
		}
	}()
}

// Legacy solution
func BFSLegacy(m *models.Maze) {
	// queue of cell positions
	q := queue.NewLinked([]int{m.Start.Pos})
	visited := make([]bool, m.N_Cells)
	parents := make([]int, m.N_Cells)

	for !q.IsEmpty() {
		pos, _ := q.Get()
		// found the maze's exit
		if pos == m.End.Pos {
			// reconstruct path
			m.Path = RebuildPath(parents, pos, m.Start.Pos)
			return
		}

		visited[pos] = true
		for _, neighbour := range m.GetNeighbours(pos) {
			if !visited[neighbour.Pos] {
				parents[neighbour.Pos] = pos
				q.Offer(neighbour.Pos)
				visited[neighbour.Pos] = true
			}
		}
	}
}

func BFS(m *models.Maze) {
	paths := make(chan []int, 1) // Creating a buffered channel
	var wg sync.WaitGroup
	BFSHelper(m, m.Start.Pos, &wg, paths)
	wg.Wait()
	close(paths) // Close the channel when all paths have finished

	m.Path = <-paths
}
