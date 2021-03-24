package tree

func recurisive_sumOfLeftLeaves(node *TreeNode,is_left bool)int{
	if node == nil{
		return 0
	}
	if node.Left == nil && node.Right == nil{
		if is_left{
			return node.Val
		}else{
			return 0
		}
	}
	return recurisive_sumOfLeftLeaves(node.Left,true) + recurisive_sumOfLeftLeaves(node.Right,false)
}

func sumOfLeftLeaves(root *TreeNode) int {
	return recurisive_sumOfLeftLeaves(root,false)
}