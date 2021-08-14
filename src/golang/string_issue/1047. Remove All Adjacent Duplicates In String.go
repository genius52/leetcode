package string_issue

import "container/list"

func removeDuplicates(s string) string {
	var l int = len(s)
	var q list.List
	for i := 0;i < l;i++{
		if q.Len() == 0{
			q.PushBack(s[i])
		}else{
			var pre uint8 = q.Back().Value.(uint8)
			if pre == s[i]{
				q.Remove(q.Back())
			}else{
				q.PushBack(s[i])
			}
		}
	}
	var res string
	for q.Len() > 0{
		res += string(q.Front().Value.(uint8))
		q.Remove(q.Front())
	}
	return res
}