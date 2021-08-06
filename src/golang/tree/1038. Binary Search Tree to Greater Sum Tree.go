package tree

func recursive_bstToGst(node* TreeNode,pre *int){
	if node == nil{
		return
	}
	recursive_bstToGst(node.Right,pre)
	node.Val += *pre
	*pre = node.Val
	recursive_bstToGst(node.Left,pre)
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	recursive_bstToGst(root,&sum)
	return root
}