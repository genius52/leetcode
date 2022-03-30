package list_queue

//Input: head = [1,0,1]
//Output: 5
func getDecimalValue(head *ListNode) int {
	var res int
	for head != nil{
		res <<= 1
		if head.Val == 1{
			res |= 1
		}
		head = head.Next
	}
	return res
}