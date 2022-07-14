package tree


func find_path(node *TreeNode,sum int) int{
	if nil == node {
		return 0
	}
	var res int = 0
	if node.Val == sum{
		res += 1
	}
	return res + find_path(node.Left,sum - node.Val) + find_path(node.Right,sum - node.Val)
}

func pathSum(root *TreeNode, sum int) int {
	if nil == root {
		return 0
	}
	return find_path(root,sum) + pathSum(root.Left,sum) + pathSum(root.Right,sum)
}