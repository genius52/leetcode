package string_issue

import (
	"container/list"
	"sort"
)

func longestWord(words []string) string {
	sort.Strings(words)
	var dict map[string]bool = make(map[string]bool)
	var l int = len(words)
	var min_len int = 2147483647
	for i := 0;i < l;i++{
		cur_len := len(words[i])
		dict[words[i]] = true
		if cur_len < min_len{
			min_len = cur_len
		}
	}
	if min_len > 1{
		return ""
	}
	var q list.List
	for i := 0;i < l;i++{
		cur_len := len(words[i])
		if cur_len == min_len{
			q.PushBack(words[i])
		}
	}
	var res string
	for q.Len() > 0{
		cur_len := q.Len()
		res = q.Front().Value.(string)
		for i := 0;i < cur_len;i++{
			var s string = q.Front().Value.(string)
			q.Remove(q.Front())
			for i := 0;i < 26;i++{
				var s2 string = s + string('a' + i)
				if _,ok := dict[s2];ok{
					q.PushBack(s2)
				}
			}
		}
	}
	return res
}