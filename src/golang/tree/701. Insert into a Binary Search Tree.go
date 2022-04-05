package tree

func recursive_insertIntoBST(node *TreeNode,val int)*TreeNode{
	if node == nil{
		var cur TreeNode
		cur.Val = val
		return &cur
	}
	if node.Val > val{
		new_left := insertIntoBST(node.Left,val)
		node.Left = new_left
	}else{
		new_right := insertIntoBST(node.Right,val)
		node.Right = new_right
	}
	return node
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		var cur TreeNode
		cur.Val = val
		return &cur
	}
	recursive_insertIntoBST(root,val)
	return root
}