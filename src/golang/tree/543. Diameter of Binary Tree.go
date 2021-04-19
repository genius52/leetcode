package tree

//para1: longest way from left leaf to root or right leaf to root
//para2: longest way from left leaf to right leaf
func dp_diameterOfBinaryTree(node *TreeNode)(int,int){
	if node == nil{
		return 0,0
	}
	l1,l2 := dp_diameterOfBinaryTree(node.Left)
	r1,r2 := dp_diameterOfBinaryTree(node.Right)
	return max_int(1 + l1,1 + r1),max_int_number(l2,r2,l1 + r1 + 1)
}

func diameterOfBinaryTree(root *TreeNode) int {
	l1,l2 := dp_diameterOfBinaryTree(root)
	return max_int(l1,l2) - 1
}