package tree

func checkTree(root *TreeNode) bool {
	if root == nil || root.Left == nil || root.Right == nil{
		return false
	}
	return root.Val == root.Left.Val + root.Right.Val
}