package list_queue

import "container/list"

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	var q1 list.List
//	var q2 list.List
//	for l1 != nil{
//		q1.PushFront(l1)
//		l1 = l1.Next
//	}
//	for l2 != nil{
//		q2.PushFront(l2)
//		l2 = l2.Next
//	}
//	var pre *ListNode = nil
//	var upgrade bool = false
//	for q1.Len() > 0 && q2.Len() > 0{
//		var n1 *ListNode = q1.Front().Value.(*ListNode)
//		var n2 *ListNode = q2.Front().Value.(*ListNode)
//		q1.Remove(q1.Front())
//		q2.Remove(q2.Front())
//		var n ListNode
//		n.Val = n1.Val + n2.Val
//		if upgrade{
//			n.Val++
//		}
//		if n.Val >= 10{
//			n.Val -= 10
//			upgrade = true
//		}else{
//			upgrade = false
//		}
//		n.Next = pre
//		pre = &n
//	}
//	for q1.Len() > 0{
//		var n1 *ListNode = q1.Front().Value.(*ListNode)
//		q1.Remove(q1.Front())
//		var n ListNode
//		n.Val = n1.Val
//		if upgrade{
//			n.Val++
//		}
//		if n.Val >= 10{
//			n.Val -= 10
//			upgrade = true
//		}else{
//			upgrade = false
//		}
//		n.Next = pre
//		pre = &n
//	}
//	for q2.Len() > 0{
//		var n2 *ListNode = q2.Front().Value.(*ListNode)
//		q2.Remove(q2.Front())
//		var n ListNode
//		n.Val = n2.Val
//		if upgrade{
//			n.Val++
//		}
//		if n.Val >= 10{
//			n.Val -= 10
//			upgrade = true
//		}else{
//			upgrade = false
//		}
//		n.Next = pre
//		pre = &n
//	}
//	if upgrade{
//		var n ListNode
//		n.Val = 1
//		n.Next = pre
//		pre = &n
//	}
//	return pre
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var q1 list.List
	var q2 list.List
	for l1 != nil{
		q1.PushFront(l1)
		l1 = l1.Next
	}
	for l2 != nil{
		q2.PushFront(l2)
		l2 = l2.Next
	}
	var pre *ListNode = nil
	var upgrade bool = false
	for q1.Len() > 0 || q2.Len() > 0 || upgrade{
		var n ListNode
		if q1.Len() > 0{
			var n1 *ListNode = q1.Front().Value.(*ListNode)
			n.Val += n1.Val
			q1.Remove(q1.Front())
		}
		if q2.Len() > 0{
			var n2 *ListNode = q2.Front().Value.(*ListNode)
			n.Val += n2.Val
			q2.Remove(q2.Front())
		}
		if upgrade{
			n.Val++
		}
		if n.Val >= 10{
			n.Val -= 10
			upgrade = true
		}else{
			upgrade = false
		}
		n.Next = pre
		pre = &n
	}
	return pre
}