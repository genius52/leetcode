package string_issue

import "container/list"

//输入：s = "abcabcababcc"
//输出：true
//解释：
//"" -> "abc" -> "abcabc" -> "abcabcabc" -> "abcabcababcc"
//因此，"abcabcababcc" 有效。


func isValid(s string) bool{
	var q list.List
	var l int = len(s)
	for i := 0;i < l;i++{
		if s[i] == 'a' || s[i] == 'b'{
			q.PushBack(s[i])
		}else if s[i] == 'c'{
			if q.Len() < 2{
				return false
			}
			var back1 uint8 = q.Back().Value.(uint8)
			q.Remove(q.Back())
			var back2 uint8 = q.Back().Value.(uint8)
			q.Remove(q.Back())
			if back1 != 'b' || back2 != 'a'{
				return false
			}
		}
	}
	return q.Len() == 0
}

//"aabbcc"
//func isValid(s string) bool {
//	var a_cnt int = 0
//	var b_cnt int = 0
//	var c_cnt int = 0
//	var l int = len(s)
//	for i := 0;i < l;i++{
//		if s[i] == 'a'{
//			a_cnt++
//		}else if s[i] == 'b'{
//			b_cnt++
//		}else if s[i] == 'c'{
//			c_cnt++
//		}
//		if c_cnt > b_cnt || b_cnt > a_cnt || c_cnt > a_cnt{
//			return false
//		}
//	}
//	if a_cnt != b_cnt || b_cnt != c_cnt || a_cnt != c_cnt{
//		return false
//	}
//	return true
//}