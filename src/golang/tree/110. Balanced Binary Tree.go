package tree

import "math"

func recursive_isBalanced(node *TreeNode)(bool,int){
	if node == nil{
		return true,0
	}
	res1,left_depth := recursive_isBalanced(node.Left)
	res2,right_depth := recursive_isBalanced(node.Right)
	if res1 && res2 && math.Abs(float64(left_depth - right_depth)) <= 1{
		return true,1 + max_int(left_depth,right_depth)
	}
	return false,0
}

func isBalanced(root *TreeNode) bool {
	res,_ := recursive_isBalanced(root)
	return res
}