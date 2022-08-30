package string_issue

import (
	"container/list"
	"strings"
)

func removeStars(s string) string {
	var l int = len(s)
	var idx int = 0
	var q list.List
	for idx < l{
		if s[idx] != '*'{
			q.PushBack(s[idx])
		}else{
			if q.Len() > 0{
				q.Remove(q.Back())
			}
		}
		idx++
	}
	var buf strings.Builder
	for q.Len() > 0{
		buf.WriteByte(q.Front().Value.(uint8))
		q.Remove(q.Front())
	}
	return buf.String()
}