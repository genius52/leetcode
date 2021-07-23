package tree

func dfs_flipMatchVoyage(node *TreeNode,voyage []int,pos *int,res *[]int)bool{
	if node == nil{
		return true
	}
	if node.Val != voyage[*pos]{
		return false
	}
	*pos++
	ret := dfs_flipMatchVoyage(node.Left,voyage,pos,res) && dfs_flipMatchVoyage(node.Right,voyage,pos,res)
	if ret {
		return true
	}
	if !ret{
		ret = dfs_flipMatchVoyage(node.Right,voyage,pos,res) && dfs_flipMatchVoyage(node.Left,voyage,pos,res)
	}
	if ret{
		//*pos++
		*res = append([]int{node.Val},*res...)
	}
	return ret
}

func FlipMatchVoyage(root *TreeNode, voyage []int) []int {
	var res []int
	var pos int = 0
	ret := dfs_flipMatchVoyage(root,voyage,&pos,&res)
	if !ret{
		return []int{-1}
	}
	return res
}