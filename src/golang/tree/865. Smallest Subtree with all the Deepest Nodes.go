package tree

func level_visit_865(node *TreeNode,cur_depth int)(max_depth int,res *TreeNode){
	if nil == node{
		return cur_depth,node
	}
	cur_depth++
	left_max_depth,leftnode := level_visit_865(node.Left,cur_depth)
	right_max_depth,rightnode := level_visit_865(node.Right,cur_depth)
	if left_max_depth == right_max_depth{
		return left_max_depth,node
	}else{
		if left_max_depth > right_max_depth{
			return left_max_depth,leftnode
		}else{
			return right_max_depth,rightnode
		}
	}
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if nil == root{
		return root
	}
	_,res :=  level_visit_865(root,0)
	return res
}
