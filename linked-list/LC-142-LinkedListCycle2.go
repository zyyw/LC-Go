package linked_list

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	m := make(map[*ListNode]bool)
	for head != nil {
		_, ok := m[head]
		if ok {
			return head
		}
		m[head] = true
		head = head.Next
	}

	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow != fast {
		return nil
	}
	fast = head
	slow = slow.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
