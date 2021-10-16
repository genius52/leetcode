package tree

func recursive_goodNodes(node *TreeNode,pre_max int)int{
	if node == nil{
		return 0
	}
	var res int = 0
	if node.Val >= pre_max{
		res++
		pre_max = node.Val
	}
	return res + recursive_goodNodes(node.Left,pre_max) + recursive_goodNodes(node.Right,pre_max)
}

func goodNodes(root *TreeNode) int {
	return recursive_goodNodes(root,-2147483648)
}