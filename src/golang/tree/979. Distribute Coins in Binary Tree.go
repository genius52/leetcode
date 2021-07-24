package tree

import "math"

func post_visit_tree(node *TreeNode,steps *int) int{
	if nil == node{
		return 0
	}

	left_val := post_visit_tree(node.Left,steps)
	right_val := post_visit_tree(node.Right,steps)
	*steps += int(math.Abs(float64(left_val + right_val + node.Val - 1)))
	return left_val + right_val + node.Val - 1
}

func distributeCoins(root *TreeNode) int {
	var steps int = 0
	post_visit_tree(root,&steps)
	return steps
}