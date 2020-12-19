package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func find_target(node *TreeNode,x int)*TreeNode{
	if node == nil{
		return nil
	}
	if node.Val == x{
		return node
	}
	left := find_target(node.Left,x)
	if left != nil{
		return left
	}
	return find_target(node.Right,x)
}

func calc_children(node *TreeNode)int{
	if node == nil{
		return 0
	}
	return  1 + calc_children(node.Left) + calc_children(node.Right)
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	if root == nil{
		return false
	}
	x_node := find_target(root,x)
	left_cnt := calc_children(x_node.Left)
	right_cnt := calc_children(x_node.Right)
	rest_cnt := n - 1 - left_cnt - right_cnt
	if left_cnt > right_cnt + rest_cnt || right_cnt > left_cnt + rest_cnt || rest_cnt > left_cnt + right_cnt{
		return true
	}
	return false
}