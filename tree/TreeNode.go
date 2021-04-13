package tree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Node struct {
	Val      int
	Children []*Node
}
