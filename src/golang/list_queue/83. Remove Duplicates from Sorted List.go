package list_queue


func DeleteDuplicates(head *ListNode) *ListNode{
	var left *ListNode = head
	var right *ListNode = head
	for left != nil{
		for right != nil && right.Val == left.Val{
			right = right.Next
		}
		left.Next = right
		left = right
	}
	return head
}

//recursive solution
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