package list_queue

func Partition(head *ListNode, x int) *ListNode {
	var small_head *ListNode = new(ListNode)
	small_head.Next = head
	var small_visit *ListNode = small_head
	var big_head *ListNode = new(ListNode)
	big_head.Next = head
	var big_visit *ListNode = big_head
	var visit *ListNode = head
	for visit != nil{
		if visit.Val < x{
			small_visit.Next = visit
			small_visit = visit
		}else{
			big_visit.Next = visit
			big_visit = visit
		}
		visit = visit.Next
	}
	big_visit.Next = nil
	small_visit.Next = big_head.Next
	return small_head.Next
}