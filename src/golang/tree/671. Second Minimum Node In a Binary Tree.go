package tree

func get_min(node *TreeNode)int{
	if node == nil{
		return 2147483647
	}
	left := get_min(node.Left)
	right := get_min(node.Right)
	return min_int_number(node.Val,left,right)
}

func get_second(node *TreeNode,min_num int)int{
	if node == nil{
		return -1
	}
	left := get_second(node.Left,min_num)
	right := get_second(node.Right,min_num)
	if left == -1{
		if right == -1{
			if node.Val == min_num{
				return -1
			}else{
				return node.Val
			}
		}else{
			if node.Val == min_num{
				return right
			}else{
				return min_int_number(node.Val,right)
			}
		}
	}else{
		if right == -1{
			if node.Val == min_num{
				return left
			}else{
				return min_int_number(node.Val,left)
			}
		}else{
			if node.Val == min_num{
				return min_int_number(left,right)
			}else{
				return min_int_number(node.Val,left,right)
			}
		}
	}
}

func FindSecondMinimumValue(root *TreeNode) int {
	if root == nil{
		return -1
	}
	min_num := get_min(root)
	return get_second(root,min_num)
}