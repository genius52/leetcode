package tree

func isUnivalTree(root *TreeNode) bool {
	if root == nil{
		return true
	}
	if root.Left != nil{
		if root.Val != root.Left.Val{
			return false
		}
	}
	if root.Right != nil{
		if root.Val != root.Right.Val{
			return false
		}
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}