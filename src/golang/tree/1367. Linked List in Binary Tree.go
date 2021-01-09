package tree

//Input: head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
//Output: true
func dp_isSubPath(cur *ListNode,root *TreeNode)bool{
	if cur == nil{
		return true
	}
	if root == nil{
		return false
	}
	return cur.Val == root.Val && (dp_isSubPath(cur.Next,root.Left) || dp_isSubPath(cur.Next,root.Right))
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil{
		return true
	}
	if root == nil{
		return false
	}
	return dp_isSubPath(head,root) || isSubPath(head,root.Left) || isSubPath(head,root.Right)
}