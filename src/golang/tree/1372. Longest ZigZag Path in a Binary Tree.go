package tree

func dp_longestZigZag(node *TreeNode)(int,int,int){
	if node == nil{
		return 0,0,0
	}
	var cur_left_max int = 0
	var cur_right_max int = 0
	var cur_disconnect_max int = 0
	if node.Left == nil && node.Right == nil{
		return 0,0,0
	}
	if node.Left != nil{
		l,r,n := dp_longestZigZag(node.Left)
		cur_left_max = 1 + r
		cur_disconnect_max = max_int_number(n,l,cur_left_max)
	}
	if node.Right != nil{
		l,r,n := dp_longestZigZag(node.Right)
		cur_right_max = 1 + l
		cur_disconnect_max = max_int_number(cur_disconnect_max,n,r,cur_right_max)
	}
	return (cur_left_max),(cur_right_max),(cur_disconnect_max)
}

func LongestZigZag(root *TreeNode) int{
	l,r,n := dp_longestZigZag(root)
	return max_int_number(l,r,n)
}

//func dp_longestZigZag(node *TreeNode,prev_left bool)int{
//	if node == nil{
//		return 0
//	}
//	if prev_left{
//		return 1 + dp_longestZigZag(node.Right,false)
//	}else{
//		return 1 + dp_longestZigZag(node.Left,true)
//	}
//}
//
//func longestZigZag(root *TreeNode) int {
//	if root == nil{
//		return 0
//	}
//	return max_int_number(dp_longestZigZag(root.Left,true),
//		dp_longestZigZag(root.Right,false),longestZigZag(root.Left),longestZigZag(root.Right))
//}