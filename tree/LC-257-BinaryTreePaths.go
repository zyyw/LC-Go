package tree

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	ret := make([]string, 0)
	binaryTreePathsHelper(root, "", &ret)
	return ret
}

func binaryTreePathsHelper(root *TreeNode, path string, ret *[]string) {
	if root == nil {
		return
	}

	path = fmt.Sprintf("%s->%d", path, root.Val)

	// check leaf node
	if root.Left == nil && root.Right == nil {
		*ret = append(*ret, path[2:])
		return
	}

	binaryTreePathsHelper(root.Left, path, ret)
	binaryTreePathsHelper(root.Right, path, ret)
}
