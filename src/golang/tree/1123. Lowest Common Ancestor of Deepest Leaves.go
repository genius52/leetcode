package tree

//dfs
func dp_lcaDeepestLeaves(node *TreeNode,level int)(int,*TreeNode){
	if node == nil{
		return level - 1,nil
	}
	if node.Left == nil && node.Right == nil{
		return level,node
	}
	left_depth,left := dp_lcaDeepestLeaves(node.Left,level + 1)
	right_depth,right := dp_lcaDeepestLeaves(node.Right,level + 1)
	if left_depth == right_depth{
		return left_depth,node
	}else if left_depth > right_depth{
		return left_depth,left
	}else{
		return right_depth,right
	}
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_,ret := dp_lcaDeepestLeaves(root,0)
	return ret
}