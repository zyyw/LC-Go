package linked_list

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// slow, fast 拆分成前后两半
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	head2 := slow.Next
	slow.Next = nil
	head2 = reverseList2(head2)

	// merge
	dummy := ListNode{
		Val:  0,
		Next: nil,
	}
	tail := &dummy
	for head2 != nil {
		tail.Next = head
		head = head.Next
		tail.Next.Next = head2
		head2 = head2.Next
		tail = tail.Next.Next
	}
	if head != nil {
		tail.Next = head
	}
}
