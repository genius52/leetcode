package tree

//Input: [4,9,0,5,1]
//    4
//   / \
//  9   0
// / \
//5   1
//Output: 1026
func visit_sumNumbers(node* TreeNode,sum int,total *int){
	if node == nil{
		return
	}
	sum = sum * 10 + node.Val
	if node.Left == nil && node.Right == nil{
		*total = *total + sum
		return
	}
	if node.Left != nil{
		visit_sumNumbers(node.Left,sum,total)
	}
	if node.Right != nil{
		visit_sumNumbers(node.Right,sum,total)
	}
}

func sumNumbers(root *TreeNode) int {
	var total int = 0
	visit_sumNumbers(root,0,&total)
	return total
}
