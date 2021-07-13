package tree

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil{
		return 0
	}
	if root.Val < low{
		return rangeSumBST(root.Right,low,high)
	}else if root.Val > high{
		return rangeSumBST(root.Left,low,high)
	}else{
		return rangeSumBST(root.Left,low,high) + root.Val + rangeSumBST(root.Right,low,high)
	}
}