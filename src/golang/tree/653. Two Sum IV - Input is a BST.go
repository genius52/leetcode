package tree

func recursive_findTarget(node *TreeNode,k int,record map[int]bool)bool{
	if node == nil{
		return false
	}
	target := k - node.Val
	if _,ok := record[target];ok{
		return true
	}
	record[node.Val] = true
	return recursive_findTarget(node.Left,k,record) || recursive_findTarget(node.Right,k,record)
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil{
		return false
	}
	var record map[int]bool = make(map[int]bool)
	return recursive_findTarget(root,k,record)
}