package list_queue

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SwapNodes(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	var left *ListNode = nil
	var right *ListNode = nil
	var i int = 0
	var visit *ListNode = head
	for visit != nil{
		if right != nil{
			right = right.Next
		}
		i++
		if i == k{
			left = visit
			right = head
		}
		visit = visit.Next
	}
	left.Val,right.Val = right.Val,left.Val
	return head
}