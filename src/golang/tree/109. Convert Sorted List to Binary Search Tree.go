package tree

//func divide_sortedListToBST(lnode *ListNode,begin int,end int)*TreeNode{
//	if begin >= end{
//		return nil
//	}
//	var mid int = begin + (end - begin)/2
//	var i int = 0
//	var visit *ListNode = lnode
//	for i != mid{
//		visit = visit.Next
//		i++
//	}
//	var tnode TreeNode
//	tnode.Val = visit.Val
//	tnode.Left = divide_sortedListToBST(lnode,begin,mid)
//	tnode.Right = divide_sortedListToBST(lnode,mid+1,end)
//	return &tnode
//}
//
//func sortedListToBST(head *ListNode) *TreeNode {
//	var l int = 0
//	var visit *ListNode = head
//	for visit != nil{
//		visit = visit.Next
//		l++
//	}
//	return divide_sortedListToBST(head,0,l)
//}

func recursive_sortedListToBST(begin *ListNode,end *ListNode)*TreeNode{
	if begin == end{
		return nil
	}
	var slow *ListNode = begin
	var fast *ListNode = begin
	for fast != end && fast.Next != end{
		slow = slow.Next
		fast = fast.Next.Next
	}
	var node *TreeNode = new(TreeNode)
	node.Val = slow.Val
	node.Left = recursive_sortedListToBST(begin,slow)
	node.Right = recursive_sortedListToBST(slow.Next,end)
	return node
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil{
		return nil
	}
	return recursive_sortedListToBST(head,nil)
}