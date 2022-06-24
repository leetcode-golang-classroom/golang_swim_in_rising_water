package sol

import "container/heap"

type Pair struct {
	Row, Col int
}
type AdjacentNode struct {
	Time  int
	Coord Pair
}
type AdjacentMinHeap []AdjacentNode

func (h *AdjacentMinHeap) Len() int {
	return len(*h)
}
func (h *AdjacentMinHeap) Less(i, j int) bool {
	return (*h)[i].Time < (*h)[j].Time
}
func (h *AdjacentMinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *AdjacentMinHeap) Push(value interface{}) {
	*h = append(*h, value.(AdjacentNode))
}
func (h *AdjacentMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func swimInWater(grid [][]int) int {
	N := len(grid)
	visit := make(map[Pair]struct{})
	directions := []Pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	priorityQueue := &AdjacentMinHeap{AdjacentNode{Time: grid[0][0], Coord: Pair{Row: 0, Col: 0}}}
	heap.Init(priorityQueue)
	for priorityQueue.Len() != 0 {
		node := heap.Pop(priorityQueue).(AdjacentNode)
		if node.Coord.Row == N-1 && node.Coord.Col == N-1 {
			return node.Time
		}
		for _, direction := range directions {
			shiftedRow := node.Coord.Row + direction.Row
			shiftedCol := node.Coord.Col + direction.Col
			if _, ok := visit[Pair{Row: shiftedRow, Col: shiftedCol}]; shiftedRow < 0 || shiftedRow == N || shiftedCol < 0 || shiftedCol == N || ok {
				continue
			}
			visit[Pair{Row: shiftedRow, Col: shiftedCol}] = struct{}{}
			heap.Push(priorityQueue,
				AdjacentNode{Time: max(node.Time, grid[shiftedRow][shiftedCol]),
					Coord: Pair{Row: shiftedRow, Col: shiftedCol}})
		}
	}
	return 0
}
