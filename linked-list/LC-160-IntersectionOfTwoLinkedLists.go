package linked_list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	m := make(map[*ListNode]bool)
	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		_, ok := m[headB]
		if ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

// 方法二：不用 map
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	h1, h2 := headA, headB
	for h1 != nil || h2 != nil {
		if h1 == h2 {
			return h1
		}
		if h1 == nil {
			h1 = headB
		} else {
			h1 = h1.Next
		}
		if h2 == nil {
			h2 = headA
		} else {
			h2 = h2.Next
		}
	}

	// h1 == nil && h2 == nil
	return nil
}
