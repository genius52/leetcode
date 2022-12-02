package list_queue

func RemoveNodes(head *ListNode) *ListNode {
	var pre *ListNode = head
	var cur *ListNode = head.Next
	pre.Next = nil
	for cur != nil{
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	var visit *ListNode = pre.Next
	var right *ListNode = pre
	var max_val int = pre.Val
	for visit != nil{
		if visit.Val >= max_val {
			max_val = visit.Val
			right.Next = visit
			right = visit
		}
		visit = visit.Next
	}
	right.Next = nil
	cur = pre.Next
	pre.Next = nil
	for cur != nil{
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}