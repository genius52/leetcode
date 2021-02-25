package list_queue


//148
//Input: 4->2->1->3
//Output: 1->2->3->4
//Input: -1->5->3->4->0
//Output: -1->0->3->4->5
func merge_sortList(l1 *ListNode,l2 *ListNode) *ListNode{
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var head *ListNode
	if l1.Val < l2.Val{
		head = l1
		l1 = l1.Next
	}else{
		head = l2
		l2 = l2.Next
	}
	var visit *ListNode = head
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			visit.Next = l1
			visit = l1
			l1 = l1.Next
		}else{
			visit.Next = l2
			visit = l2
			l2 = l2.Next
		}
	}
	if l1 == nil{
		visit.Next = l2
	}else if l2 == nil{
		visit.Next = l1
	}
	return head
}
//-1->5->3->4->0
func sortList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next{
		return head
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next
		if fast != nil{
			fast = fast.Next
		}
	}
	head2 := slow.Next
	slow.Next = nil
	newhead1 := sortList(head)
	newhead2 := sortList(head2)
	return merge_sortList(newhead1,newhead2)
}
