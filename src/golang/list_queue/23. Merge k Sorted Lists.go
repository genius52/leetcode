package list_queue

/**
 * Definition for singly-linked list_queue.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge_sort(list1 *ListNode,list2 *ListNode)*ListNode{
	if list1 == nil{
		return list2
	}
	if list2 == nil{
		return list1
	}
	var head *ListNode = nil
	var visit *ListNode = nil
	for list1 != nil && list2 != nil{
		if list1.Val < list2.Val{
			if head == nil{
				head = list1
				visit = list1
			}else{
				visit.Next = list1
				visit = list1
			}
			list1 = list1.Next
		}else{
			if head == nil{
				head = list2
				visit = list2
			}else{
				visit.Next = list2
				visit = list2
			}
			list2 = list2.Next
		}
	}
	for list1 != nil{
		visit.Next = list1
		visit = list1
		list1 = list1.Next
	}
	for list2 != nil{
		visit.Next = list2
		visit = list2
		list2 = list2.Next
	}
	return head
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0{
		return nil
	}
	var prev_lists []*ListNode = make([]*ListNode,len(lists))
	copy(prev_lists,lists)
	for len(prev_lists) > 1{
		var lists_merge []*ListNode
		l := len(prev_lists)
		for i := 0;i < l - 1;i += 2{
			new_list := merge_sort(prev_lists[i],prev_lists[i+1])
			lists_merge = append(lists_merge, new_list)
		}
		if l % 2 == 1{
			lists_merge = append(lists_merge,prev_lists[l - 1])
		}
		prev_lists = lists_merge
	}
	return prev_lists[0]
}