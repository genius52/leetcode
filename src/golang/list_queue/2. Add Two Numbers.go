package list_queue

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil
	var pre *ListNode = nil
	var next_val int = 0
	for l1 != nil || l2 != nil || next_val != 0{
		var cur *ListNode = new(ListNode)
		if l1 != nil && l2 != nil{
			cur.Val = l1.Val + l2.Val + next_val
			next_val = cur.Val / 10
			cur.Val %= 10
			l1 = l1.Next
			l2 = l2.Next
		}else if l1 != nil{
			cur.Val = l1.Val + next_val
			next_val = cur.Val / 10
			cur.Val %= 10
			l1 = l1.Next
		}else if l2 != nil{
			cur.Val = l2.Val + next_val
			next_val = cur.Val / 10
			cur.Val %= 10
			l2 = l2.Next
		}else{
			cur.Val = next_val
			next_val = 0
		}
		if head == nil{
			head = cur
			pre = cur
		}else{
			pre.Next = cur
			pre = cur
		}
	}
	return head
}