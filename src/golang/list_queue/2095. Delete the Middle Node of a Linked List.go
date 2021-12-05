package list_queue

func deleteMiddle(head *ListNode) *ListNode {
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
	var c int = 0
	for visit != nil {
		if c == mid {
			next := visit.Next
			visit.Next = next.Next
		}
		visit = visit.Next
		c++
	}
	return head
}
