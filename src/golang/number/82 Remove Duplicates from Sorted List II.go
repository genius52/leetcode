package number

import "sort"

//Definition for singly-linked list_queue.
type ListNode struct {
	Val int
	Next *ListNode
}

//82
//Input: 1->1->1->2->3
//Output: 2->3
func deleteDuplicates(head *ListNode) *ListNode {
	var record map[int]int = make(map[int]int)
	for head != nil{
		if _,ok := record[head.Val];ok{
			record[head.Val]++
		}else{
			record[head.Val] = 1
		}
		head = head.Next
	}
	var keys []int
	for k := range record {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var ret *ListNode = nil
	var visit *ListNode = nil
	for _,n := range keys{
		if record[n] > 1{
			continue
		}
		var node ListNode
		node.Val = n
		if ret == nil{
			ret = &node
			visit = &node
		}else{
			visit.Next = &node
			visit = visit.Next
		}
	}
	return ret
}