package list_queue

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//Input: head = [1,2,-3,3,1]
//Output: [3,1]
//[1,3,2,-3,-2,5,5,-5,1]
func RemoveZeroSumSublists(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	var record map[int]*ListNode = make(map[int]*ListNode)
	var cur *ListNode = head
	var prefix int = 0
	for cur != nil{
		prefix += cur.Val
		record[prefix] = cur
		cur = cur.Next
	}
	var dummy *ListNode = new(ListNode)
	dummy.Next = head
	cur = dummy
	prefix = 0
	for cur != nil{
		prefix += cur.Val
		if node,ok := record[prefix];ok{
			if node != cur{
				cur.Next = node.Next
				//delete(record,prefix)
			}
		}
		cur = cur.Next
	}
	cur = dummy.Next
	for cur != nil{
		if cur.Val == 0{
			cur = cur.Next
		}else{
			break
		}
	}
	return cur
}