package linked_list

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := getLength(head)
	k = k % n
	if k == 0 {
		return head
	}

	prev, node := head, head
	for k > 0 {
		node = node.Next
		k--
	}
	for node.Next != nil {
		node = node.Next
		prev = prev.Next
	}

	newHead := prev.Next
	prev.Next = nil
	node.Next = head

	return newHead
}

// get length of ListNode
func getLength(head *ListNode) int {
	n := 0
	for head != nil {
		n++
		head = head.Next
	}
	return n
}
