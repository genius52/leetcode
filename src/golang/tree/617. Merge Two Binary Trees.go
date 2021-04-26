package tree

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil{
		return root2
	}
	if root2 == nil{
		return root1
	}
	root1.Val += root2.Val
	new_left := mergeTrees(root1.Left,root2.Left)
	new_right := mergeTrees(root1.Right,root2.Right)
	root1.Left = new_left
	root1.Right = new_right
	return root1
}