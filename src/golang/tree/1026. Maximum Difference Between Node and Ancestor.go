package tree

import "math"

func pre_visit(node *TreeNode,min int,max int)int{
	if nil == node{
		return 0
	}
	diff := int(math.Max(math.Abs(float64(node.Val - min)),math.Abs(float64(node.Val - max))))
	min = int(math.Min(float64(node.Val),float64(min)))
	max = int(math.Max(float64(node.Val),float64(max)))
	return max_int_number(diff,pre_visit(node.Left,min,max),pre_visit(node.Right,min,max))
}

func maxAncestorDiff(root *TreeNode) int {
	if nil == root{
		return 0
	}
	return pre_visit(root,root.Val,root.Val)
}