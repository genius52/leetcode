package tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	pre_left := root.Left
	pre_right := root.Right
	root.Right = pre_left
	root.Left = pre_right
	invertTree(pre_left)
	invertTree(pre_right)
	return root
}