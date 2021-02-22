package tree

func pre_faltten(node *TreeNode)*TreeNode{
	if nil == node{
		return nil
	}
	pre_right := node.Right
	left_tail := pre_faltten(node.Left)
	if nil != left_tail{
		left_tail.Right = pre_right
	}
	if nil != node.Left{
		node.Right = node.Left
	}
	node.Left = nil
	right_tail := pre_faltten(pre_right)
	if nil != right_tail{
		return right_tail
	}
	if nil != left_tail{
		return left_tail
	}
	return node
}

func flatten(root *TreeNode)  {
	pre_faltten(root)
}
