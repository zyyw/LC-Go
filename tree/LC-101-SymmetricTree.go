package tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetricHelper(root.Left, root.Right)
}

func symmetricHelper(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return symmetricHelper(p.Left, q.Right) && symmetricHelper(p.Right, q.Left)
}
