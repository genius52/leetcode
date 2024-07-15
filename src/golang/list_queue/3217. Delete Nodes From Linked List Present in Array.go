package list_queue

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ModifiedList(nums []int, head *ListNode) *ListNode {
	var record map[int]bool = make(map[int]bool)
	for _, n := range nums {
		record[n] = true
	}
	var new_head *ListNode = nil
	var pre *ListNode = nil
	var visit *ListNode = head
	for visit != nil {
		if _, ok := record[visit.Val]; !ok { //可以保留
			if new_head == nil {
				new_head = visit
			}
			if pre == nil {
				pre = visit
			} else {
				pre.Next = visit
				pre = visit
			}
		}
		visit = visit.Next
	}
	if pre != nil {
		pre.Next = nil
	}
	return new_head
}
