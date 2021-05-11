package tree


func dp_longestUnivaluePath(node *TreeNode,pre_val int)(int,int){//same_val_len,diff_val_len
	if node == nil{
		return 0,0
	}
	left_same_len,left_diff_len := dp_longestUnivaluePath(node.Left,node.Val)
	right_same_len,right_diff_len :=  dp_longestUnivaluePath(node.Right,node.Val)
	if pre_val == node.Val{
		return 1 + max_int_number(left_same_len,right_same_len),max_int_number(left_same_len + right_same_len,left_diff_len,right_diff_len)
	}else{
		return 0,max_int_number(left_same_len + right_same_len,left_diff_len,right_diff_len)
	}
}

func longestUnivaluePath(root *TreeNode) int {
	val1,val2 := dp_longestUnivaluePath(root,2147483647)
	return max_int_number(val1,val2)
}