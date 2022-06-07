package list_queue

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var visit1 *ListNode = headA
	var visit2 *ListNode = headB
	var l1 int = 0
	var l2 int = 0
	for visit1 != nil{
		visit1 = visit1.Next
		l1++
	}
	for visit2 != nil{
		visit2 = visit2.Next
		l2++
	}
	var diff int = l1 - l2
	if diff > 0{
		for diff > 0{
			headA = headA.Next
			diff--
		}
	}else if diff < 0{
		for diff < 0{
			headB = headB.Next
			diff++
		}
	}
	for headA != nil && headB != nil{
		if headA == headB{
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode{
	var p1 *ListNode = headA
	var p2 *ListNode = headB
	for p1 != nil && p2 != nil && p1 != p2{
		p1 = p1.Next
		p2 = p2.Next
		if p1 == p2{
			return p1
		}
		if p1 == nil{
			p1 = headB
		}
		if p2 == nil{
			p2 = headA
		}
	}
	return p1
}