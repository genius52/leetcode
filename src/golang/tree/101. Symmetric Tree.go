package tree

func dfs_isSymmetric(left *TreeNode,right *TreeNode)bool{
	if left == nil && right == nil{
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil){
		return false
	}
	if left.Val != right.Val{
		return false
	}
	return dfs_isSymmetric(left.Left,right.Right) && dfs_isSymmetric(left.Right,right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return dfs_isSymmetric(root.Left,root.Right)
}