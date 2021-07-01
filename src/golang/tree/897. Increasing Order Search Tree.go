package tree

var pre_node *TreeNode = nil
func inorder_increasingBST(node *TreeNode,new_root **TreeNode){
	if node == nil{
		return
	}
	inorder_increasingBST(node.Left,new_root)
	var new_node TreeNode
	new_node.Val = node.Val
	if pre_node != nil{
		pre_node.Right = &new_node
	}
	pre_node = &new_node
	if *new_root == nil{
		*new_root = &new_node
	}
	inorder_increasingBST(node.Right,new_root)
}

func IncreasingBST(root *TreeNode) *TreeNode {
	var new_root *TreeNode
	pre_node = nil
	inorder_increasingBST(root,&new_root)
	return new_root
}