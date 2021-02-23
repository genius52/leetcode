package tree

import "math"

//124
//return1: open way max
//return2: closed way max
func post_visit(node *TreeNode)(int,int){
	if node == nil{
		return math.MinInt32,math.MinInt32
	}
	if node.Left == nil && node.Right == nil{
		return node.Val,node.Val
	}
	left_root_leaf_sum,left_leaf_leaf_sum := post_visit(node.Left)
	right_root_leaf_sum,right_leaf_leaf_sum := post_visit(node.Right)
	return max_int_number(left_root_leaf_sum + node.Val,right_root_leaf_sum + node.Val,node.Val),
		max_int_number(left_root_leaf_sum,left_leaf_leaf_sum,right_root_leaf_sum,right_leaf_leaf_sum,left_root_leaf_sum + right_root_leaf_sum + node.Val)
}

func maxPathSum(root *TreeNode) int {
	r1,r2 := post_visit(root)
	return max_int_number(r1,r2)
}