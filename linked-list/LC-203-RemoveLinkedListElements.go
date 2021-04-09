package linked_list

func removeElements(head *ListNode, val int) *ListNode {
	dummy := ListNode{
		Val:  0,
		Next: nil,
	}
	tail := &dummy
	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			// 这个节点的值不等于 val，需要保留
			tail.Next = head
			head = head.Next
			tail = tail.Next
		}
	}
	tail.Next = nil

	return dummy.Next
}
