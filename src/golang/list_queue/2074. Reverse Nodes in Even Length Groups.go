package list_queue

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	var total_len int = 0
	visit := head
	for visit != nil{
		visit = visit.Next
		total_len++
	}
	var cur_len int = 0
	target_len := 1
	last := head
	cur := head
	new_head := head
	for cur != nil{
		if target_len | 1 == target_len{
			for i := 0;i < target_len;i++{
				last = cur
				cur = cur.Next
				cur_len++
			}
		}else{
			var prev *ListNode = cur
			var visit *ListNode = cur.Next
			var next *ListNode = cur.Next.Next
			var next_group *ListNode = nil
			var reverse_head *ListNode = nil
			var reverse_tail *ListNode = nil
			for i := 0;i < target_len - 1;i++{
				if i == 0{
					reverse_tail = prev
				}
				visit.Next = prev
				if i == target_len - 2{
					next_group = next
					reverse_head = visit
				}
				prev = visit
				visit = next
				if next != nil{
					next = next.Next
				}
			}
			cur_len += target_len
			last.Next = reverse_head
			reverse_tail.Next = next_group
			cur = next_group
			last = reverse_tail
		}
		target_len++
		rest_len := total_len - cur_len
		if target_len > rest_len{
			target_len = rest_len
		}
	}
	return new_head
}