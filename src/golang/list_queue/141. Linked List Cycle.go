package list_queue

func hasCycle(head *ListNode) bool {
	var slow *ListNode = head
	var fast *ListNode = head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			return true
		}
	}
	return false
}