package linked_list

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := ListNode{
		Val:  0,
		Next: head,
	}
	for head != nil && head.Next != nil {
		if head.Val != head.Next.Val {
			// no dup
			head = head.Next
		} else {
			// head.Next is a dup of head
			head.Next = head.Next.Next
		}
	}

	return dummy.Next
}
