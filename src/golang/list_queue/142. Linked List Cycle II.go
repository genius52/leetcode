package list_queue

//Input: head = [3,2,0,-4], pos = 1
//Output: tail connects to node index 1
func detectCycle(head *ListNode) *ListNode {
	var slow *ListNode = head
	var fast *ListNode = head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			break
		}
	}
	if fast == nil || fast.Next == nil{
		return nil
	}
	var steps int = 1
	slow = slow.Next
	for slow != fast{
		slow = slow.Next
		steps++
	}
	var visit *ListNode = head
	for steps > 0{
		visit = visit.Next
		steps--
	}
	var res *ListNode = head
	for res != visit{
		res = res.Next
		visit = visit.Next
	}
	return res
}