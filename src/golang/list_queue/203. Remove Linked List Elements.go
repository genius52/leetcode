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

func removeElements2(head *ListNode, val int) *ListNode {
	var left *ListNode = head
	var new_head *ListNode = nil
	for left != nil{
		for left != nil && left.Val == val{
			left = left.Next
		}
		if left == nil{
			break
		}
		if new_head == nil{
			new_head = left
		}
		var right *ListNode = left.Next
		for right != nil && right.Val == val{
			right = right.Next
		}
		left.Next = right
		left = right
	}
	return new_head
}
