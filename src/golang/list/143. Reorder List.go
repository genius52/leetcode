package list

import "container/list"

type ListNode struct {
	Val int
	Next *ListNode
}

//Given a singly linked list L: L0→L1→…→Ln-1→Ln,
//reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
func ReorderList(head *ListNode)  {
	if head == nil || head.Next == nil{
		return
	}
	var list1 list.List
	var list2 list.List
	slow := head
	fast := head
	for fast != nil && fast.Next != nil{
		list1.PushBack(slow)
		slow = slow.Next
		fast = fast.Next.Next
	}
	for slow != nil{
		list2.PushBack(slow)
		slow = slow.Next
	}
	var new_head *ListNode = nil
	var visit *ListNode = nil
	for list1.Len() > 0 && list2.Len() > 0{
		n1 := list1.Front().Value.(*ListNode)
		n2 := list2.Back().Value.(*ListNode)
		n1.Next = n2
		n2.Next = nil
		if new_head == nil{
			new_head = n1
			visit = n2
		}else{
			visit.Next = n1
			visit = n2
		}
		list1.Remove(list1.Front())
		list2.Remove(list2.Back())
	}
	if list2.Len() > 0{
		n := list2.Back().Value.(*ListNode)
		visit.Next = n
		n.Next = nil
	}
	head = new_head
}