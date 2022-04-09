package list_queue

import "sort"

//Definition for singly-linked list_queue.
//type ListNode struct {
//	Val int
//	Next *ListNode
//}

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

//Input: head = [1,2,3,3,4,4,5]
//Output: [1,2,5]
func DeleteDuplicates2(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	var slow *ListNode = head
	var fast *ListNode = head
	for fast.Next != nil{
		if fast.Next.Val == slow.Val{
			fast = fast.Next
		}else{
			break
		}
	}
	if slow == fast{
		next := DeleteDuplicates2(fast.Next)
		slow.Next = next
		return slow
	}else{
		next := DeleteDuplicates2(fast.Next)
		return next
	}
}

func deleteDuplicates2(head *ListNode) *ListNode{
	var left *ListNode = head
	var pre *ListNode
	var new_head *ListNode = nil
	for left != nil{
		var right *ListNode = left.Next
		for right != nil && right.Val == left.Val{
			right = right.Next
		}
		if right == left.Next{
			if new_head == nil{
				new_head = left
			}
			if pre == nil{
				pre = left
			}else{
				pre.Next = left
				pre = left
			}
		}
		left = right
	}
	return new_head
}