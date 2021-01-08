package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sum_maxProduct(node *TreeNode)int{
	if node == nil{
		return 0
	}
	return node.Val + sum_maxProduct(node.Left) + sum_maxProduct(node.Right)
}

func dp_maxProduct(node *TreeNode,total int)(int,int){
	if node == nil{
		return 0,0
	}
	left_sum,left_product := dp_maxProduct(node.Left,total)
	right_sum,right_product := dp_maxProduct(node.Right,total)
	cur_sum := node.Val + left_sum + right_sum
	cur_product := (total - cur_sum) * cur_sum
	max_product := max_int_number(cur_product,left_product,right_product)
	return cur_sum,max_product
}

func maxProduct(root *TreeNode) int {
	var total int = sum_maxProduct(root)
	_,res := dp_maxProduct(root,total)
	return res % 1000000007
}