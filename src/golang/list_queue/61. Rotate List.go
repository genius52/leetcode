package list_queue

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0{
		return head
	}
	var l int = 0
	var visit *ListNode = head
	var end *ListNode = visit
	for visit != nil{
		l++
		end = visit
		visit = visit.Next
	}
	k %= l
	if k == 0{
		return head
	}
	visit = head
	k = l - k
	var new_head *ListNode = nil
	for k > 1{
		visit = visit.Next
		k--
	}
	new_head = visit.Next
	visit.Next = nil
	end.Next = head
	return new_head
}