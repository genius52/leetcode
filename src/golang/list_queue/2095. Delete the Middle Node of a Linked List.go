package list_queue

func DeleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var l int = 0
	var visit *ListNode = head
	for visit != nil {
		visit = visit.Next
		l++
	}
	var mid int = l/2 - 1
	visit = head
	for visit != nil {
		if mid == 0 {
			next := visit.Next
			visit.Next = next.Next
			break
		}
		visit = visit.Next
		mid--
	}
	return head
}
