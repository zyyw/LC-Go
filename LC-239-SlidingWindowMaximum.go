package LC_Go

import "container/heap"

type ValueIndex struct {
	val, index int
}

type MaxHeapValueIndex []ValueIndex

func (h MaxHeapValueIndex) Len() int {
	return len(h)
}

func (h MaxHeapValueIndex) Less(i, j int) bool {
	if h[i].val == h[j].val {
		return h[i].index < h[j].index
	}
	return h[i].val > h[j].val
}

func (h MaxHeapValueIndex) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeapValueIndex) Push(x interface{}) {
	item := x.(ValueIndex)
	*h = append(*h, item)
}

func (h *MaxHeapValueIndex) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	ret := make([]int, 0)
	if k <= 0 || k > len(nums) {
		return ret
	}

	// maxHeap
	maxHeap := make(MaxHeapValueIndex, 0)
	for i := 0; i < k; i++ {
		maxHeap = append(maxHeap, ValueIndex{
			val:   nums[i],
			index: i,
		})
	}
	heap.Init(&maxHeap)

	ret = append(ret, maxHeap[0].val)
	for i := k; i < len(nums); i++ {
		heap.Push(&maxHeap, ValueIndex{
			val:   nums[i],
			index: i,
		})
		for maxHeap[0].index < i-k+1 {
			heap.Pop(&maxHeap)
		}
		ret = append(ret, maxHeap[0].val)
	}

	return ret
}
