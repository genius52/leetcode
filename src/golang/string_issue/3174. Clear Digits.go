package string_issue

import "container/list"

func clearDigits(s string) string {
	var q list.List
	for _, c := range s {
		if c >= '0' && c <= '9' {
			q.Remove(q.Back())
		} else {
			q.PushBack(c)
		}
	}
	var res string
	for q.Len() > 0 {
		res += string(q.Front().Value.(int32))
		q.Remove(q.Front())
	}
	return res
}
