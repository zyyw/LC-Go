package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := ListNode{
		Val:  0,
		Next: head,
	}
	prev := &dummy
	for n > 1 {
		head = head.Next
		n--
		if head == nil {
			// n > 原链表的长度
			return nil
		}
	}
	// prev 和 head 同时往后移动，直到 head 抵达最后一个节点
	for head.Next != nil {
		head = head.Next
		prev = prev.Next
	}
	// 移除 prev.Next 这个节点
	prev.Next = prev.Next.Next

	return dummy.Next
}
