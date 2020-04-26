package tree

func recursive_countNodes(node *TreeNode,num int)int{
	var left_max int
	var right_max int
	if node.Right != nil{
		right_max = recursive_countNodes(node.Right,num * 2 + 1)
	}
	if node.Left != nil{
		left_max = recursive_countNodes(node.Left,num * 2)
	}
	return max_int_number(num,left_max,right_max)
}

func countNodes(root *TreeNode) int {
	if root == nil{
		return 0
	}
	return recursive_countNodes(root,1)
}