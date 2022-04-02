package list_queue

func recursive_reverseList(node *ListNode,new_head **ListNode)*ListNode{
	if node.Next == nil{
		*new_head = node
		return node
	}
	new_last := recursive_reverseList(node.Next,new_head)
	node.Next = nil
	new_last.Next = node
	return node
}

func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	var new_head *ListNode
	recursive_reverseList(head,&new_head)
	return new_head
}

func reverseList2(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	var new_head *ListNode = nil
	var pre *ListNode = head
	var cur *ListNode = head.Next
	pre.Next = nil
	for cur != nil{
		next := cur.Next
		if next == nil{
			new_head = cur
		}
		cur.Next = pre
		pre = cur
		cur = next
	}
	return new_head
}