package list_queue

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	var visit *ListNode = head
	for visit != nil{
		if visit.Val == head.Val{
			visit = visit.Next
		}else{
			break
		}
	}
	if visit == nil{
		head.Next = nil
	}else{
		next := deleteDuplicates1(visit)
		head.Next = next
	}
	return head
}