package LC_Go

import "container/heap"

type NumItem struct {
	Num, Freq int
}

// implement MinHeap
type MinHeapNum []NumItem

func (h MinHeapNum) Len() int {
	return len(h)
}

func (h MinHeapNum) Less(i, j int) bool {
	return h[i].Freq < h[j].Freq
}

func (h MinHeapNum) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeapNum) Push(x interface{}) {
	item := x.(NumItem)
	*h = append(*h, item)
}

func (h *MinHeapNum) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	// 1. 统计数字频率 (num, freq)
	m := make(map[int]int)
	for _, num := range nums {
		freq := m[num]
		m[num] = freq + 1
	}

	// 2. 找出数字频率 top K
	// initialize minHeap
	minHeap := make(MinHeapNum, k)
	for num, freq := range m {
		if len(minHeap) < k {
			// initializing minHeap
			minHeap = append(minHeap, NumItem{
				Num:  num,
				Freq: freq,
			})
			if len(minHeap) == k {
				heap.Init(&minHeap)
			}
		} else {
			// 与堆顶元素进行比较
			if minHeap[0].Freq < freq {
				heap.Pop(&minHeap)
				heap.Push(&minHeap, NumItem{
					Num:  num,
					Freq: freq,
				})
			}
		}
	}

	// 3. return
	ret := make([]int, len(minHeap))
	for i := 0; i < len(minHeap); i++ {
		ret[i] = minHeap[i].Num
	}
	return ret
}
