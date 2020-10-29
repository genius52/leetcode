package list_queue

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
	if head == nil{
		return 0
	}
	var record []bool = make([]bool,10000)
	for _,v := range G{
		record[v] = true
	}
	//var begin *ListNode = head
	var res int = 0
	var visit *ListNode = head
	for visit != nil {
		//search connected start
		for visit != nil && !record[visit.Val]{
			visit = visit.Next
		}
		if visit != nil{
			res++
		}
		//search not connected point
		for visit != nil && record[visit.Val]{
			visit = visit.Next
		}
	}
	return res
}