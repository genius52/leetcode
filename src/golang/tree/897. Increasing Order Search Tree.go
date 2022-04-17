package tree


func inorder_increasingBST(node *TreeNode,pre_node **TreeNode,new_root **TreeNode){
	if node == nil{
		return
	}
	inorder_increasingBST(node.Left,pre_node,new_root)
	node.Left = nil
	if *pre_node != nil{
		(*pre_node).Right = node
	}
	*pre_node = node
	if *new_root == nil{
		*new_root = *pre_node
	}
	tmp := node.Right
	node.Right = nil
	inorder_increasingBST(tmp,pre_node,new_root)
}

func IncreasingBST(root *TreeNode) *TreeNode {
	var new_root *TreeNode
	var pre_node *TreeNode = nil
	inorder_increasingBST(root,&pre_node,&new_root)
	return new_root
}