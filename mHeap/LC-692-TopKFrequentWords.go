package mHeap

import "container/heap"

type WordItem struct {
	Word string
	Freq int
}

type MinHeapWord []WordItem

func (h MinHeapWord) Len() int {
	return len(h)
}

func (h MinHeapWord) Less(i, j int) bool {
	if h[i].Freq == h[j].Freq {
		return h[i].Word > h[j].Word
	}
	return h[i].Freq < h[j].Freq
}

func (h MinHeapWord) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeapWord) Push(x interface{}) {
	word := x.(WordItem)
	*h = append(*h, word)
}

func (h *MinHeapWord) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func topKFrequentWords(words []string, k int) []string {
	if len(words) == 0 {
		return []string{}
	}

	// 1. 统计词频
	m := make(map[string]int)
	for _, word := range words {
		freq := m[word]
		m[word] = freq + 1
	}

	// 2. 根据词频，找 top K
	minHeap := make(MinHeapWord, 0)
	for word, freq := range m {
		if len(minHeap) < k {
			// initializing minHeap
			minHeap = append(minHeap, WordItem{
				Word: word,
				Freq: freq,
			})
			if len(minHeap) == k {
				heap.Init(&minHeap)
			}
		} else {
			// 与堆顶元素比较
			if minHeap[0].Freq < freq || (minHeap[0].Freq == freq && word < minHeap[0].Word) {
				heap.Pop(&minHeap)
				heap.Push(&minHeap, WordItem{
					Word: word,
					Freq: freq,
				})
			}
		}
	}

	// 3. return
	ret := make([]string, len(minHeap))
	idx := len(minHeap) - 1
	for len(minHeap) > 0 {
		item := heap.Pop(&minHeap).(WordItem)
		ret[idx] = item.Word // 降序
		idx--
	}
	return ret
}
