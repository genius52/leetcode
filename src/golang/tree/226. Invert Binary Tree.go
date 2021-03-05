package tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	old_left := root.Left
	old_right := root.Right
	root.Right = invertTree(old_left)
	root.Left = invertTree(old_right)
	return root
}