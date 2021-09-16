package tree

func dfs_sumEvenGrandparent(node *TreeNode,parent_val int,grand_val int)int{
	if nil == node{
		return 0
	}
	var cur_sum int = 0
	if grand_val % 2 == 0 {
		cur_sum = node.Val
	}
	return cur_sum + dfs_sumEvenGrandparent(node.Left,node.Val,parent_val) + dfs_sumEvenGrandparent(node.Right,node.Val,parent_val)
}

func sumEvenGrandparent(root *TreeNode) int {
	if nil == root{
		return 0
	}
	return dfs_sumEvenGrandparent(root.Left,root.Val,1) + dfs_sumEvenGrandparent(root.Right,root.Val,1)
}