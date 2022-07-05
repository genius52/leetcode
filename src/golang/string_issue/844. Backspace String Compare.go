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

func BackspaceCompare(s string, t string) bool{
	var l1 int = len(s)
	var l2 int = len(t)
	var i int = l1 - 1
	var j int = l2 - 1
	for i >= 0 || j >= 0{
		var star1 int = 0
		for i >= 0 && (s[i] == '#' || star1 > 0){
			if s[i] == '#'{
				star1++
			}else{
				star1--
			}
			i--
		}
		var star2 int = 0
		for j >= 0 && (t[j] == '#' || star2 > 0){
			if t[j] == '#'{
				star2++
			}else{
				star2--
			}
			j--
		}
		if i < 0 || j < 0{
			break
		}
		if i >= 0 && j >= 0 &&  s[i] != t[j]{
			return false
		}
		i--
		j--
	}
	return i < 0 && j < 0
}

//func BackspaceCompare(s string, t string) bool{
//	var l1 int = len(s)
//	var l2 int = len(t)
//	var idx1 int = l1 - 1
//	var idx2 int = l2 - 1
//	var cnt1 int = 0
//	var cnt2 int = 0
//	for idx1 >= 0 || idx2 >= 0{
//		for idx1 >= 0 && s[idx1] == '#'{
//			cnt1++
//			idx1--
//		}
//		for cnt1 > 0 {
//			idx1--
//			cnt1--
//		}
//		if idx1 > 0 && s[idx1] == '#'{
//			continue
//		}
//
//		for idx2 >= 0 && t[idx2] == '#'{
//			cnt2++
//			idx2--
//		}
//		for cnt2 > 0{
//			idx2--
//			cnt2--
//		}
//		if idx2 > 0 && t[idx2] == '#'{
//			continue
//		}
//		for idx1 >= 0 && s[idx1] != '#' && idx2 >= 0 && t[idx2] != '#'{
//			if s[idx1] != t[idx2]{
//				return false
//			}
//			idx1--
//			idx2--
//		}
//	}
//	return idx1 <= -1 && idx2 <= -1
//}