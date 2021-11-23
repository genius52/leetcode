package list_queue

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	visit1 := list1
	var first *ListNode
	var cnt int = a
	for b >= 0{
		if cnt == 1{
			first = visit1
		}
		visit1 = visit1.Next
		cnt--
		b--
	}
	first.Next = list2
	visit2 := list2
	for visit2.Next != nil{
		visit2 = visit2.Next
	}
	visit2.Next = visit1
	return list1
}