package tree

func recursive_sufficientSubset(node *TreeNode,limit int,sum int)*TreeNode{
	if node == nil{
		return node
	}
	if node.Left == nil && node.Right == nil{
		if (sum + node.Val) < limit{
			return nil
		}
		return node
	}
	node.Left = recursive_sufficientSubset(node.Left,limit,sum + node.Val)
	node.Right = recursive_sufficientSubset(node.Right,limit,sum + node.Val)
	if node.Left == nil && node.Right == nil{
		return nil
	}
	return node
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	return recursive_sufficientSubset(root,limit,0)
}