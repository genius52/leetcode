package list_queue

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return true
	}
	//reverse half
	var slow *ListNode = head
	var fast *ListNode = head
	var pre *ListNode = nil
	var next *ListNode = head.Next
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow.Next = pre
		pre = slow
		slow = next
		next = next.Next
	}
	if fast != nil{
		slow = slow.Next
	}
	for pre != nil && slow != nil{
		if pre.Val != slow.Val{
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}
	return true
}