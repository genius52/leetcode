package list_queue

import (
	"container/list"
)

type ListNode struct {
	Val int
	Next *ListNode
}

//Given a singly linked list_queue L: L0→L1→…→Ln-1→Ln,
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


func reverse(node *ListNode,new_head **ListNode)*ListNode{
	if node == nil{
		return nil
	}
	next := reverse(node.Next,new_head)
	node.Next = nil
	if next != nil{
		next.Next = node
		return node
	}else{
		*new_head = node
		return *new_head
	}
}

func merge(node1 *ListNode,node2 *ListNode){
	var n1 *ListNode = node1
	var n2 *ListNode = node2
	for n1 != nil && n2 != nil{
		tmp1 := n1.Next
		tmp2 := n2.Next
		n1.Next = n2
		n1 = tmp1
		n2.Next = n1
		n2 = tmp2
	}
	if n1 != nil{
		n1.Next = nil
	}
	if n2 != nil{
		n2.Next = nil
	}
}

func ReorderList2(head *ListNode) {
	if head == nil || head.Next == nil{
		return
	}
	slow := head
	fast := head
	var pre_end *ListNode = nil
	for fast != nil && fast.Next != nil{ //find mid
		pre_end = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil{
		slow = slow.Next
		pre_end = pre_end.Next
	}
	pre_end.Next = nil
	var reverse_start *ListNode = slow
	var new_head *ListNode
	reverse(reverse_start,&new_head)
	var visit *ListNode = head
	merge(visit,new_head)
}