package list_queue

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse_doubleIt(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var new_head *ListNode = nil
	var pre *ListNode = head
	var cur *ListNode = head.Next
	pre.Next = nil
	for cur != nil {
		next := cur.Next
		if next == nil {
			new_head = cur
		}
		cur.Next = pre
		pre = cur
		cur = next
	}
	return new_head
}

func doubleIt(head *ListNode) *ListNode {
	new_head := reverse_doubleIt(head)
	var upgrade bool = false
	var visit *ListNode = new_head
	var last *ListNode = nil
	for visit != nil {
		val := visit.Val * 2
		if upgrade {
			val++
		}
		if val >= 10 {
			upgrade = true
		} else {
			upgrade = false
		}
		visit.Val = val % 10
		last = visit
		visit = visit.Next
	}
	if upgrade {
		add := new(ListNode)
		add.Val = 1
		last.Next = add
	}
	return reverse_doubleIt(new_head)
}
