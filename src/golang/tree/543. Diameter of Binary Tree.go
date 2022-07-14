package tree

func recursive_diameterOfBinaryTree(node *TreeNode,res *int)int{
	if node == nil{
		return 0
	}
	left := recursive_diameterOfBinaryTree(node.Left,res)
	right := recursive_diameterOfBinaryTree(node.Right,res)
	cur := left + right
	if cur > *res{
		*res = cur
	}
	return max_int(left,right) + 1
}

func DiameterOfBinaryTree(root *TreeNode) int {
	var res int = 0
	recursive_diameterOfBinaryTree(root,&res)
	return res
}

//para1: longest way from left leaf to root or right leaf to root
//para2: longest way from left leaf to right leaf
//func dp_diameterOfBinaryTree(node *TreeNode)(int,int){
//	if node == nil{
//		return 0,0
//	}
//	l1,l2 := dp_diameterOfBinaryTree(node.Left)
//	r1,r2 := dp_diameterOfBinaryTree(node.Right)
//	return max_int(1 + l1,1 + r1),max_int_number(l2,r2,l1 + r1 + 1)
//}
//
//func diameterOfBinaryTree(root *TreeNode) int {
//	l1,l2 := dp_diameterOfBinaryTree(root)
//	return max_int(l1,l2) - 1
//}
