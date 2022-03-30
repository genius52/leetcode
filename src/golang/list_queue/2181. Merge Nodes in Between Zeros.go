package list_queue

//[0,3,1,0,4,5,2,0]
func mergeNodes(head *ListNode) *ListNode {
	var new_head *ListNode = head.Next
	var left *ListNode = head.Next
	var right *ListNode = head.Next
	var sum int = 0
	for right != nil{
		if right.Val != 0{
			sum += right.Val
			right = right.Next
		}else{
			left.Val = sum
			sum = 0
			left.Next = right.Next
			left = right.Next
			right = right.Next
		}
	}
	return new_head
}