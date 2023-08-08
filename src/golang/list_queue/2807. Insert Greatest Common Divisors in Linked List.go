package list_queue

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var cur *ListNode = head
	var next *ListNode = head.Next
	for next != nil {
		var add *ListNode = new(ListNode)
		add.Val = gcd(cur.Val, next.Val)
		cur.Next = add
		add.Next = next
		cur = next
		next = next.Next
	}
	return head
}
