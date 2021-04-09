package linked_list

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	carry, sum := 0, 0
	base := 10
	dummy := ListNode{
		Val:  0,
		Next: nil,
	}
	tail := &dummy

	for l1 != nil && l2 != nil {
		sum = carry + l1.Val + l2.Val
		tail.Next = &ListNode{
			Val:  sum % base,
			Next: nil,
		}
		tail = tail.Next
		carry = sum / base
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum = carry + l1.Val
		tail.Next = &ListNode{
			Val:  sum % base,
			Next: nil,
		}
		tail = tail.Next
		carry = sum / base
		l1 = l1.Next
	}
	for l2 != nil {
		sum = carry + l2.Val
		tail.Next = &ListNode{
			Val:  sum % base,
			Next: nil,
		}
		tail = tail.Next
		carry = sum / base
		l2 = l2.Next
	}
	if carry > 0 {
		tail.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return dummy.Next
}
