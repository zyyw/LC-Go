package linked_list

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head2 := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return head2
}

// 迭代法
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := ListNode{
		Val:  0,
		Next: head,
	}
	for head.Next != nil {
		node := head.Next
		head.Next = node.Next
		node.Next = dummy.Next
		dummy.Next = node
	}

	return dummy.Next
}
