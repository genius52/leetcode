package list_queue

import "container/list"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//Input: [2,7,4,3,5]
//Output: [7,0,5,5,0]
func NextLargerNodes(head *ListNode) []int {
	var stack list.List
	var nodes []int
	for head != nil{
		nodes = append(nodes,head.Val)
		head = head.Next
	}
	var l int = len(nodes)
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		if stack.Len() == 0{
			stack.PushBack(i)
		}else{
			prev_index := stack.Back().Value.(int)
			for stack.Len() > 0 && nodes[prev_index] < nodes[i]{
				res[prev_index] = nodes[i]
				stack.Remove(stack.Back())
				if stack.Len() > 0{
					prev_index = stack.Back().Value.(int)
				}
			}
			stack.PushBack(i)
		}
	}
	return res
}