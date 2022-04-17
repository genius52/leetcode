package tree

func preorder_convertBST(node *TreeNode,pre_sum *int){
	if node == nil{
		return
	}
	preorder_convertBST(node.Right,pre_sum)
	node.Val += *pre_sum
	*pre_sum = node.Val
	preorder_convertBST(node.Left,pre_sum)
}

func ConvertBST(root *TreeNode) *TreeNode {
	var pre_sum int = 0
	preorder_convertBST(root,&pre_sum)
	return root
}