package list

//ugly solution!!!
//Input: -1->5->3->4->0
//Output: -1->0->3->4->5
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