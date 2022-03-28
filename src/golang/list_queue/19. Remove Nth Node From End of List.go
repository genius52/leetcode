package list_queue

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var cnt int = 0
	var right *ListNode = head
	for cnt < n && right != nil{
		right = right.Next
		cnt++
	}
	if right == nil{
		return head.Next
	}
	var left *ListNode = head
	for right.Next != nil{
		right = right.Next
		left = left.Next
	}
	left.Next = left.Next.Next
	return head
}