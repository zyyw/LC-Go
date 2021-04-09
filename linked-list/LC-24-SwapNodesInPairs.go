package linked_list

func swapPairs(head *ListNode) *ListNode {
	dummy := ListNode{
		Val:  0,
		Next: head,
	}
	tail := &dummy

	// o -> 1 -> 2 -> 3
	// t    h    n
	for head != nil && head.Next != nil {
		node := head.Next
		head.Next = node.Next
		node.Next = head
		tail.Next = node
		tail = head
		head = head.Next
	}

	return dummy.Next
}
