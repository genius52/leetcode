package tree

func postorder_averageOfSubtree(node *TreeNode,res *int)(int,int){
	if node == nil{
		return 0,0
	}
	left_sum,left_cnt := postorder_averageOfSubtree(node.Left,res)
	right_sum,right_cnt := postorder_averageOfSubtree(node.Right,res)
	sum := left_sum + right_sum + node.Val
	aver := sum/(left_cnt + right_cnt + 1)
	if aver == node.Val{
		*res++
	}
	return sum,left_cnt + right_cnt + 1
}

func averageOfSubtree(root *TreeNode) int {
	var res int
	postorder_averageOfSubtree(root,&res)
	return res
}