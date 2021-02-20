package list_queue

func reverse_list(begin *ListNode,end *ListNode)*ListNode{
	if begin == end{
		begin.Next = nil
		return end
	}
	ret := reverse_list(begin.Next,end)
	begin.Next = nil
	ret.Next = begin
	return begin
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right{
		return head
	}
	var begin *ListNode = head
	var dummy *ListNode = new(ListNode)
	dummy.Next = head
	var prev *ListNode = dummy
	right = right - left + 1
	for begin != nil{
		left--
		if left == 0{
			break
		}else{
			begin = begin.Next
			prev = prev.Next
		}
	}
	if begin == nil{
		return head
	}
	var end *ListNode = begin
	for end != nil{
		right--
		if right == 0{
			break
		}else{
			end = end.Next
		}
	}
	if end == nil{
		return head
	}
	next := end.Next
	reverse_list(begin,end)
	prev.Next = end
	begin.Next = next
	return dummy.Next
}