package tree

import "math"

func recursive_isBalanced(node *TreeNode)int{
	if node == nil{
		return 0
	}
	left_depth := recursive_isBalanced(node.Left)
	right_depth := recursive_isBalanced(node.Right)
	if left_depth != -1 && right_depth != -1 && math.Abs(float64(left_depth - right_depth)) <= 1{
		return 1 + max_int(left_depth,right_depth)
	}
	return -1
}

func isBalanced(root *TreeNode) bool {
	return  recursive_isBalanced(root) != -1
}