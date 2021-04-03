package mHeap

import (
	"container/heap"
	"math"
)

type MinHeapNum []int

func (h MinHeapNum) Len() int {
	return len(h)
}

func (h MinHeapNum) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeapNum) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeapNum) Push(x interface{}) {
	item := x.(int)
	*h = append(*h, item)
}

func (h *MinHeapNum) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return math.MaxInt32
	}

	// 初始化 heap
	minHeap := make(MinHeapNum, 0)

	// 遍历
	for _, num := range nums {
		if len(minHeap) < k {
			minHeap = append(minHeap, num)
			if len(minHeap) == k {
				// 建堆
				heap.Init(&minHeap)
			}
		} else {
			if num > minHeap[0] {
				heap.Pop(&minHeap)
				heap.Push(&minHeap, num)
			}
		}
	}

	// 返回值
	return minHeap[0]
}

//// 快排序思想
func findKthLargest2(nums []int, k int) int {
	if len(nums) < k {
		return math.MaxInt32
	}

	start, end := 0, len(nums) - 1
	for start <= end {
		mi := partition(nums, start, end)
		if mi == k - 1 {
			return nums[mi]
		} else if mi < k - 1 {
			start = mi + 1
		} else {
			end = mi - 1
		}
	}
	return math.MaxInt32
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	mi := start
	for i := start+1; i <= end; i++ {
		if nums[i] > pivot {
			mi++
			// swap(nums, mi, i)
			nums[mi], nums[i] = nums[i], nums[mi]
		}
	}
	// swap(nums, start, mi)
	nums[start], nums[mi] = nums[mi], nums[start]


	// return
	return mi
}