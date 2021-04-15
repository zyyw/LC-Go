package tree

/*
	二叉树，p, q 均在给定的树中。
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	retLeft := lowestCommonAncestor(root.Left, p, q)
	retRight := lowestCommonAncestor(root.Right, p, q)

	if retLeft != nil {
		if retRight != nil {
			return root
		}
		return retLeft
	}
	return retRight
}
