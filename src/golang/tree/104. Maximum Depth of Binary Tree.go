package tree

func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	return 1 + max_int(maxDepth(root.Left),maxDepth(root.Right))
}