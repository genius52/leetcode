package list_queue

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	var even *ListNode = head
	var odd *ListNode = head.Next
	var even_tail *ListNode = head
	var odd_head *ListNode = head.Next
	for even != nil && odd != nil{
		var next_even *ListNode = nil
		var next_odd *ListNode = nil
		if odd.Next != nil{
			next_even = odd.Next
		}
		if next_even != nil{
			next_odd = next_even.Next
		}
		if next_even != nil{
			even_tail = next_even
		}
		even.Next = next_even
		even = next_even
		odd.Next = next_odd
		odd = next_odd
	}
	even_tail.Next = odd_head
	return head
}