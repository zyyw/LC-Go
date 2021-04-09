package linked_list

// 借助 reverse linked list 1: 迭代法的思想
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	d := right - left
	dummy := ListNode{
		Val:  0,
		Next: head,
	}
	prev := &dummy
	for left > 1 {
		if prev == nil {
			return nil
		}
		prev = prev.Next
		left--
	}

	// reverse list node between [left, right], 即 [prev.Next, d + 1 个]
	// left = 2, right = 4, d = 2
	// 0  -> 1  ->  2  ->  3  ->  4 -> 5
	//      prev   head   node  right
	head = prev.Next
	for d > 0 {
		node := head.Next
		head.Next = node.Next
		node.Next = prev.Next
		prev.Next = node
		d--
	}

	return dummy.Next
}
