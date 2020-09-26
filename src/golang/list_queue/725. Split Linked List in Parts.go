package list_queue

func SplitListToParts(root *ListNode, k int) []*ListNode {
	if k == 0{
		return []*ListNode{}
	}
	var l int = 0
	visit := root
	for visit != nil{
		visit = visit.Next
		l++
	}
	rest := l % k
	sub_len := l / k
	var res []*ListNode
	visit = root
	for visit != nil{
		res = append(res,visit)
		var cur_len = sub_len
		if rest > 0{
			cur_len += 1
			rest--
		}
		for i := 0;i < cur_len - 1 && visit != nil;i++{
			visit = visit.Next
		}
		if visit == nil{
			break
		}
		tmp := visit
		visit = visit.Next
		tmp.Next = nil
	}
	k -= len(res)
	for k > 0{
		res = append(res,nil)
		k--
	}
	return res
}