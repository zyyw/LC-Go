package LC_Go

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *MinHeap) Push(x interface{}) {
	node := x.(*ListNode)
	*h = append(*h, node)
}

func (h *MinHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	tail := dummy

	// initialize minHeap
	minHeap := make(MinHeap, 0)
	for _, node := range lists {
		if node != nil {
			minHeap = append(minHeap, node)
		}
	}
	heap.Init(&minHeap)

	// iterate minHeap
	for len(minHeap) > 0 {
		node := heap.Pop(&minHeap).(*ListNode)
		tail.Next = node
		tail = tail.Next
		// populating minHeap
		if node.Next != nil {
			heap.Push(&minHeap, node.Next)
		}
	}

	// return
	return dummy.Next
}
