package number

import "container/list"

func validStrings(n int) []string {
	var q list.List
	q.PushBack("0")
	q.PushBack("1")
	for n > 0 {
		var cur_len int = q.Len()
		for i := 0; i < cur_len; i++ {
			cur := q.Front().Value.(string)
			q.Remove(q.Front())
			if cur[len(cur)-1] == '1' {
				q.PushBack(cur + "0")
			}
			q.PushBack(cur + "1")
		}
		n--
	}
	var res []string
	for q.Len() > 0 {
		res = append(res, q.Front().Value.(string))
		q.Remove(q.Front())
	}
	return res
}
