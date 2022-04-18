package tree

func inorder_kthSmallest(node *TreeNode,k *int,res *int){
	if node == nil{
		return
	}
	inorder_kthSmallest(node.Left,k,res)
	if *k == 0{
		return
	}
	*k--
	if *k == 0{
		*res = node.Val
		return
	}
	inorder_kthSmallest(node.Right,k,res)
}

func kthSmallest(root *TreeNode, k int) int {
	var res int = 0
	inorder_kthSmallest(root,&k,&res)
	return res
}