package string_issue

import (
	"container/list"
	"math"
)

func makeGood(s string) string{
	var q list.List
	var l int = len(s)
	var diff int = 'a' - 'A'
	for i := 0;i < l;i++{
		if q.Len() == 0{
			q.PushBack(s[i])
		}else{
			top := q.Back().Value.(uint8)
			if int(math.Abs(float64(s[i]) - float64(top))) == diff{
				q.Remove(q.Back())
			}else{
				q.PushBack(s[i])
			}
		}
	}
	var res string
	for q.Len() != 0{
		res += string(q.Front().Value.(uint8))
		q.Remove(q.Front())
	}
	return res
}

func MakeGood(s string) string {
	var l int = len(s)
	if l == 0{
		return s
	}
	var diff int = 'a' - 'A'
	var stack []uint8 = []uint8{ s[0]}
	var pos = 1
	for pos < l{
		if(len(stack) == 0){
			stack = append(stack,s[pos])

		}else{
			last := stack[len(stack) - 1]
			if(int(math.Abs(float64(last) - float64(s[pos]))) == diff){
				stack = stack[:len(stack) - 1]
			}else{
				stack = append(stack,s[pos])
			}
		}
		pos++
	}
	var res string
	for _,c := range stack{
		res += string(c)
	}
	return res
}