package string_issue

import (
	"container/list"
	"strings"
)

func RobotWithString(s string) string {
	var l int = len(s)
	var min_char []uint8 = make([]uint8,l + 1)//min_char[i]: i之后会出现的最小字符
	min_char[l] = 'z' + 1
	for i := l - 1;i >= 0;i--{
		if s[i] < min_char[i + 1]{
			min_char[i] = s[i]
		}else{
			min_char[i] = min_char[i + 1]
		}
	}
	var q list.List
	var ss strings.Builder
	for i := 0;i < l;i++{
		q.PushBack(s[i])
		for q.Len() > 0 && q.Back().Value.(uint8) <= min_char[i + 1]{
			ss.WriteByte(q.Back().Value.(uint8))
			q.Remove(q.Back())
		}
	}
	for q.Len() > 0{
		ss.WriteByte(q.Back().Value.(uint8))
		q.Remove(q.Back())
	}
	return ss.String()
}