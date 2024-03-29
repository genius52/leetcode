package list_queue

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//Input: head = [1,2,3,4]
//Output: [2,1,4,3]
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	ret := swapPairs(head.Next.Next)
	new_head := head.Next
	new_head.Next = head
	head.Next = ret
	return new_head
}

func swapPairs2(head *ListNode) *ListNode{
	var visit *ListNode = head
	for visit != nil && visit.Next != nil{
		visit.Val,visit.Next.Val = visit.Next.Val,visit.Val
		visit = visit.Next.Next
	}
	return head
}