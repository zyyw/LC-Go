package linked_list

func deleteNode(node *ListNode) {
	for node.Next.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}

	// node.Next.Next == nil, node.Next 是尾节点
	node.Val = node.Next.Val
	node.Next = nil
}
