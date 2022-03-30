package list_queue

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil{
		return list2
	}
	if list2 == nil{
		return list1
	}
	var visit *ListNode = nil
	var head *ListNode
	if list1.Val <= list2.Val{
		head = list1
		visit = list1
		list1 = list1.Next
	}else{
		head = list2
		visit = list2
		list2 = list2.Next
	}
	for list1 != nil && list2 != nil{
		if list1.Val <= list2.Val{
			visit.Next = list1
			visit = list1
			list1 = list1.Next
		}else{
			visit.Next = list2
			visit = list2
			list2 = list2.Next
		}
	}
	if list1 != nil{
		visit.Next = list1
	}
	if list2 != nil{
		visit.Next = list2
	}
	return head
}