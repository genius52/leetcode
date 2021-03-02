package list_queue

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil{
		return nil
	}
	next := removeElements(head.Next,val)
	if head.Val == val{
		return next
	}else{
		head.Next = next
		return head
	}
}
