package tree

func recursive_lowestCommonAncestor(node *TreeNode,low int,high int)*TreeNode{
	if node == nil{
		return nil
	}
	if node.Val >= low && node.Val <= high{
		return node
	}
	if node.Val < low{
		ret := recursive_lowestCommonAncestor(node.Right,low,high)
		if ret != nil{
			return ret
		}
	}
	if node.Val > high{
		ret := recursive_lowestCommonAncestor(node.Left,low,high)
		if ret != nil{
			return ret
		}
	}
	return node
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	var low int = 0
	var high int = 0
	if q.Val < p.Val{
		low = q.Val
		high = p.Val
	}else{
		low = p.Val
		high = q.Val
	}
	return recursive_lowestCommonAncestor(root,low,high)
}