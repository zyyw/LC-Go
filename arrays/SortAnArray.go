package arrays

import (
	"container/heap"
)

/*
  problem statement:
  给你一个整数数组 nums，请你将该数组升序排列。

  思路：
	1. 选择排序，冒泡排序 - 时间复杂度：O(N^2), 空间复杂度: O(1)
	2. 快速排序 - 时间复杂度: O(N log N)，空间复杂度: O(log N) 栈递归调用
	   归并排序 - 时间复杂度: O(N log N), 空间复杂度: O(N)
	3. 堆排序
*/

// 快速排序
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	sortArrayHelper(nums, 0, len(nums)-1)
	return nums
}

func sortArrayHelper(nums []int, start, end int) {
	if start >= end {
		return
	}

	mi := partition(nums, start, end)
	sortArrayHelper(nums, start, mi-1)
	sortArrayHelper(nums, mi+1, end)
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	mi := start
	for i := start + 1; i <= end; i++ {
		if nums[i] < pivot {
			mi++
			nums[mi], nums[i] = nums[i], nums[mi] // swap(nums, mi, i)
		}
	}
	nums[start], nums[mi] = nums[mi], nums[start]
	return mi
}

// 归并排序
func sortArray2(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mi := len(nums) / 2
	left, right := make([]int, mi), make([]int, len(nums)-mi)
	for i := 0; i < mi; i++ {
		left[i] = nums[i]
	}
	for i := mi; i < len(nums); i++ {
		right[i-mi] = nums[i]
	}
	sortedLeft := sortArray2(left)
	sortedRight := sortArray2(right)
	res := merge(sortedLeft, sortedRight)
	return res
}

func merge(left, right []int) []int {
	i, j := 0, 0
	nums := make([]int, len(left)+len(right))
	k := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		nums[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		nums[k] = right[j]
		j++
		k++
	}
	return nums
}

// 堆排序
type MinHeapNums []int

func (h MinHeapNums) Len() int {
	return len(h)
}
func (h MinHeapNums) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeapNums) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeapNums) Push(x interface{}) {
	item := x.(int)
	*h = append(*h, item)
}
func (h *MinHeapNums) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func sortArray3(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	// min heap
	minHeap := make(MinHeapNums, len(nums))
	for i := 0; i < len(nums); i++ {
		minHeap[i] = nums[i]
	}
	heap.Init(&minHeap)

	idx := 0
	for len(minHeap) > 0 {
		res := heap.Pop(&minHeap).(int)
		nums[idx] = res
		idx++
	}
	return nums
}

// 冒泡排序
func sortArray4(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1] // swap(nums, j-1, j)
			}
		}
	}
	return nums
}
