package string_issue

import "container/list"

type char_cnt struct {
	c uint8
	cnt int
}

func RemoveDuplicates2(s string, k int) string {
	var q list.List
	var l int = len(s)
	for i := 0;i < l;i++{
		if q.Len() == 0{
			var obj char_cnt
			obj.c = s[i]
			obj.cnt = 1
			q.PushBack(obj)
		}else{
			last := q.Back().Value.(char_cnt)
			if last.c == s[i]{
				last.cnt++
				q.Remove(q.Back())
				if last.cnt != k{
					q.PushBack(last)
				}
			}else{
				var cur char_cnt
				cur.c = s[i]
				cur.cnt = 1
				q.PushBack(cur)
			}
		}
	}
	var res string
	for q.Len() > 0{
		cur := q.Front().Value.(char_cnt)
		for cur.cnt > 0{
			res += string(cur.c)
			cur.cnt--
		}
		q.Remove(q.Front())
	}
	return res
}