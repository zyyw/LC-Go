package LC_Go

import (
	"container/heap"
	"math"
)

type MatrixItem struct {
	X, Y int // row, col index
	Val int  // matrix[x][y]
}

type MinHeapMatrix []MatrixItem

func (h MinHeapMatrix) Len() int {
	return len(h)
}

func (h MinHeapMatrix) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeapMatrix) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeapMatrix) Push(x interface{}) {
	item := x.(MatrixItem)
	*h = append(*h, item)
}

func (h *MinHeapMatrix) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return math.MaxInt32
	}

	m, n := len(matrix), len(matrix[0])
	// visited
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	// 初始化 heap
	minHeap := make(MinHeapMatrix, 0)
	minHeap = append(minHeap, MatrixItem{
		X:   0,
		Y:   0,
		Val: matrix[0][0],
	})
	heap.Init(&minHeap)
	visited[0][0] = true

	// 遍历
	dx := [2]int{0, 1}
	dy := [2]int{1, 0}
	for k > 1 {
		cur := heap.Pop(&minHeap).(MatrixItem)
		for idx := 0; idx < 2; idx++ {
			x := cur.X + dx[idx]
			y := cur.Y + dy[idx]
			// 越界
			if !(x < m && y < n) {
				continue
			}
			// 已经被遍历过
			if visited[x][y] {
				continue
			}
			heap.Push(&minHeap, MatrixItem{
				X:   x,
				Y:   y,
				Val: matrix[x][y],
			})
			visited[x][y] = true
		}
		k--
	}

	return minHeap[0].Val
}
