package list_queue

func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil{
		fast = fast.Next
		if fast != nil{
			fast = fast.Next
		}
		slow = slow.Next
	}
	return slow
}