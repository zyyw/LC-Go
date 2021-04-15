package tree

/*
	二叉搜索树，p, q 均在给定的树中
*/

func lowestCommonAncestorOfBST(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorOfBST(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorOfBST(root.Right, p, q)
	}
	return root
}
