package tree

var max_res int = 0
func dfs_maxSumBST(node *TreeNode)(int,int,int){
	if node == nil{
		return 0,-2147483647,2147483647
	}
	// node.val is the min return value,node.Val - 1 is allowed to "left_max < node.Val && node.Val < right_min"
	var left_sum,left_min,left_max = 0,node.Val,node.Val - 1
	var right_sum,right_min,right_max = 0,node.Val + 1,node.Val
	if node.Left != nil{
		left_sum,left_min,left_max = dfs_maxSumBST(node.Left)
	}
	if node.Right != nil{
		right_sum,right_min,right_max = dfs_maxSumBST(node.Right)
	}
	if left_max < node.Val && node.Val < right_min {
		max_res = max_int(max_res,left_sum + node.Val + right_sum)
		return left_sum + node.Val + right_sum,left_min,right_max
	}
	return 0,-2147483647,2147483647
}

func MaxSumBST(root *TreeNode) int {
	if root == nil{
		return 0
	}
	max_res = 0
	dfs_maxSumBST(root)
	return max_res
}