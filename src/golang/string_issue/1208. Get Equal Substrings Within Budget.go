package string_issue

import "math"

//Input: s = "abcd", t = "bcdf", maxCost = 3
//Output: 3
//Explanation: "abc" of s can change to "bcd". That costs 3, so the maximum length is 3.
func EqualSubstring(s string, t string, maxCost int) int {
	var begin int = 0
	var end int = 0
	var l int = len(s)
	var res int = 0
	var cur_diff int = 0
	for end < l{
		if cur_diff > maxCost{
			cur_diff -= int(math.Abs(float64(s[begin]) - float64(t[begin])))
			begin++
		}else{
			cur_len := end - begin
			if cur_len > res{
				res = cur_len
			}
			cur_diff += int(math.Abs(float64(s[end]) - float64(t[end])))
			end++
		}
	}
	if cur_diff <= maxCost{
		cur_len := end - begin
		if cur_len > res{
			res = cur_len
		}
	}
	return res
}


//func dp_equalSubstring(s string,t string,pos int,cost int,is_start bool)int{
//	if pos >= len(s){
//		return 0
//	}
//	cur_cost := math.Abs(float64(s[pos]) - float64(t[pos]))
//	if is_start{
//		if cur_cost > float64(cost){
//			return 0
//		}else{
//			return 1 + dp_equalSubstring(s,t,pos + 1,cost - int(cur_cost),is_start)
//		}
//	}else{
//		var do_start int = 0
//		if cur_cost > float64(cost){
//			do_start = 0
//		}else{
//			do_start = 1 + dp_equalSubstring(s,t,pos + 1,cost - int(cur_cost),true)
//		}
//		var do_not_start int = dp_equalSubstring(s,t,pos + 1,cost,is_start)
//		if do_start > do_not_start{
//			return do_start
//		}else{
//			return do_not_start
//		}
//	}
//}
//
//func EqualSubstring(s string, t string, maxCost int) int {
//	return dp_equalSubstring(s,t,0,maxCost,false)
//}