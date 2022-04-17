package tree

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil{
		return nil
	}
	left := trimBST(root.Left,low,high)
	right := trimBST(root.Right,low,high)
	if root.Val < low || root.Val > high{
		if left != nil{
			return left
		}else{
			return right
		}
	}else{
		root.Left = left
		root.Right = right
		return root
	}
}