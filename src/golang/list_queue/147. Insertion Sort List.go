package list_queue

//ugly solution!!!
//Input: -1->5->3->4->0
//Output: -1->0->3->4->5
func InsertionSortList2(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	var dummy *ListNode = new(ListNode)
	dummy.Val = -2147483648
	dummy.Next = head
	var sort_end *ListNode = head
	for sort_end.Next != nil{
		var to_sort *ListNode = sort_end.Next
		if to_sort.Val > sort_end.Val{
			sort_end = to_sort
		}else{
			var visit *ListNode = dummy
			var pre *ListNode = dummy
			for visit.Val <= to_sort.Val{
				pre = visit
				visit = visit.Next
			}
			sort_end.Next = to_sort.Next
			pre.Next = to_sort
			to_sort.Next = visit
		}
	}
	return dummy.Next
}

func InsertionSortList(head *ListNode) *ListNode {
	var new_head *ListNode = nil
	var visit *ListNode = head
	for visit != nil{
		if new_head == nil{
			new_head = visit
			visit = visit.Next
			new_head.Next = nil
			continue
		}
		search_insert := new_head
		var last *ListNode = nil
		for search_insert != nil{
			if visit.Val < search_insert.Val{
				tmp := visit.Next
				visit.Next = search_insert
				if last != nil{
					last.Next = visit
				}else{
					new_head = visit
				}
				visit = tmp
				break
			}else{
				last = search_insert
				search_insert = search_insert.Next
			}
		}
		if search_insert == nil{
			last.Next = visit
			tmp := visit.Next
			visit.Next = search_insert
			visit = tmp
		}
	}
	return new_head
}