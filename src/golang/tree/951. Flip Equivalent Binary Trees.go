package tree

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if nil == root1 && nil == root2{
		return true
	}
	if nil == root1 || nil == root2{
		return false
	}
	if root1.Val != root2.Val{
		return false
	}
	return (flipEquiv(root1.Left,root2.Left) && flipEquiv(root1.Right,root2.Right)) ||
		(flipEquiv(root1.Left,root2.Right) && flipEquiv(root1.Right,root2.Left))
}