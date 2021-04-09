package linked_list

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := ListNode{
		Val:  0,
		Next: head,
	}
	prev := &dummy
	for head != nil && head.Next != nil {
		for head.Val == head.Next.Val {
			head = head.Next
			if head.Next == nil {
				break
			}
		}
		if prev.Next == head {
			// no dup
			prev = head
			head = head.Next
		} else {
			// contain dup [prev.Next, head]
			prev.Next = head.Next
			head = head.Next
		}
	}

	return dummy.Next
}
