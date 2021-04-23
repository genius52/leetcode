package tree

import "math"

func dfs_findTilt(node *TreeNode,total *int) int{
	if node == nil{
		return 0
	}
	left_sum := dfs_findTilt(node.Left,total)
	right_sum := dfs_findTilt(node.Right,total)
	var sum int = left_sum + right_sum + node.Val
	node.Val = int(math.Abs(float64(left_sum - right_sum)))
	*total += node.Val
	return sum
}

func findTilt(root *TreeNode) int {
	var total int = 0
	dfs_findTilt(root,&total)
	return total
}