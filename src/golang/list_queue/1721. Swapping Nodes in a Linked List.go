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
	var node1 *ListNode = nil
	var node2 *ListNode = nil
	var i int = 0
	var visit *ListNode = head
	for visit != nil{
		if node2 != nil{
			node2 = node2.Next
		}
		i++
		if i == k{
			node1 = visit
			node2 = head
		}

		visit = visit.Next
	}
	node1.Val,node2.Val = node2.Val,node1.Val
	return head
}