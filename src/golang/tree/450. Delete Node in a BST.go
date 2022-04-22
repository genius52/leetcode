package tree


func find_max_inleft(node *TreeNode)int {
	for node.Right != nil{
		node = node.Right
	}
	return node.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil{
		return root
	}
	if key == root.Val{
		if root.Left == nil{
			return root.Right
		}
		if root.Right == nil{
			return root.Left
		}
		val := find_max_inleft(root.Left)
		root.Val = val
		root.Left = deleteNode(root.Left,val)
	}else if key < root.Val{
		root.Left = deleteNode(root.Left,key)
	}else if key > root.Val{
		root.Right = deleteNode(root.Right,key)
	}
	return root
}