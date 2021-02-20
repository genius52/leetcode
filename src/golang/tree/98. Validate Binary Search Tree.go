package tree

func dfs_isValidBST(node *TreeNode,low int64,high int64)bool{
	if node == nil{
		return true
	}
	if int64(node.Val) <= low || int64(node.Val) >= high{
		return false
	}
	return dfs_isValidBST(node.Left,low,int64(node.Val)) && dfs_isValidBST(node.Right,int64(node.Val),high)
}

func isValidBST(root *TreeNode) bool {
	return dfs_isValidBST(root.Left,-(1<<63),int64(root.Val)) && dfs_isValidBST(root.Right,int64(root.Val),1<<63 - 1)
}