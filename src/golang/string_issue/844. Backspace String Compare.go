package string_issue

import "container/list"

//Input: s = "a##c", t = "#a#c"
//Output: true
func backspaceCompare(s string, t string) bool {
	var q1 list.List
	var q2 list.List
	var l1 int = len(s)
	var l2 int = len(t)
	for i := 0;i < l1;i++{
		if s[i] == '#'{
			if q1.Len() != 0{
				q1.Remove(q1.Back())
			}
		}else{
			q1.PushBack(s[i])
		}
	}
	for i := 0;i < l2;i++{
		if t[i] == '#'{
			if q2.Len() != 0{
				q2.Remove(q2.Back())
			}
		}else{
			q2.PushBack(t[i])
		}
	}
	if q1.Len() != q2.Len(){
		return false
	}
	var l3 int = q1.Len()
	for i := 0;i < l3;i++{
		if q1.Front().Value != q2.Front().Value{
			return false
		}
		q1.Remove(q1.Front())
		q2.Remove(q2.Front())
	}
	return true
}