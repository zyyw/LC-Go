package linked_list

func oddEvenList(head *ListNode) *ListNode {
	dummy1 := ListNode{
		Val:  0,
		Next: nil,
	}
	dummy2 := ListNode{
		Val:  0,
		Next: nil,
	}
	tail1, tail2 := &dummy1, &dummy2

	for head != nil && head.Next != nil {
		tail1.Next = head
		tail1 = tail1.Next
		tail2.Next = head.Next
		tail2 = tail2.Next
		head = head.Next.Next
	}
	if head != nil {
		// 链表里有奇数个元素
		tail1.Next = head
		tail1 = tail1.Next
	}
	tail1.Next = dummy2.Next
	tail2.Next = nil

	return dummy1.Next
}
