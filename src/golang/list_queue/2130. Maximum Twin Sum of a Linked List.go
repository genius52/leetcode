package list_queue

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func PairSum(head *ListNode) int {
	var l int = 0
	var visit *ListNode = head
	for visit != nil {
		visit = visit.Next
		l++
	}
	var record []int = make([]int, l/2)
	visit = head
	cnt := 0
	var res int = 0
	for visit != nil {
		if cnt < l/2 {
			record[cnt] = visit.Val
		} else {
			cur := visit.Val + record[l-1-cnt]
			if cur > res {
				res = cur
			}
		}
		visit = visit.Next
		cnt++
	}
	return res
}
