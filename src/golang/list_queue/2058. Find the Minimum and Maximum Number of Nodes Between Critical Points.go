package list_queue

func nodesBetweenCriticalPoints(head *ListNode) []int {
	if head == nil || head.Next == nil{
		return []int{-1,-1}
	}
	pre := head
	visit := head.Next
	next := head.Next.Next
	var idx int = 1
	var record []int
	for next != nil{
		if (visit.Val > pre.Val && visit.Val > next.Val) || (visit.Val < pre.Val && visit.Val < next.Val){
			record = append(record,idx)
		}
		idx++
		pre = visit
		visit = next
		next = next.Next
	}
	var l int = len(record)
	if l <= 1{
		return []int{-1,-1}
	}else{
		var min_diff int = 2147483647
		var max_diff int = record[l - 1] - record[0]
		for i := 1;i < l;i++{
			cur_diff := record[i] - record[i - 1]
			if cur_diff < min_diff{
				min_diff = cur_diff
			}
		}
		return []int{min_diff,max_diff}
	}
}