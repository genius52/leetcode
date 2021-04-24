package tree

func isSame(t1 *TreeNode, t2 *TreeNode) bool{
	if nil == t1 && nil == t2{
		return true
	}
	if nil == t1 || nil == t2{
		return false
	}
	if t1.Val != t2.Val{
		return false
	}
	return isSame(t1.Left,t2.Left) && isSame(t1.Right,t2.Right);
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if nil == t{
		return true
	}
	if nil == s{
		return false
	}
	if isSame(s,t){
		return true
	}
	return isSubtree(s.Left,t) || isSubtree(s.Right,t);
}