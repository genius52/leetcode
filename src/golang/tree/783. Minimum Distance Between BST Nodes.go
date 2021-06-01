package tree

var pre *TreeNode
func inorder_minDiffInBST(node *TreeNode)int{
	if node == nil{
		return 2147483647
	}
	left_diff := inorder_minDiffInBST(node.Left)
	var diff int = 2147483647
	if pre != nil{
		diff = node.Val - pre.Val
	}
	pre = node
	right_diff := inorder_minDiffInBST(node.Right)
	return min_int_number(left_diff,diff,right_diff)
}

func MinDiffInBST(root *TreeNode) int {
	pre = nil
	return inorder_minDiffInBST(root)
}