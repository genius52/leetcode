package string_issue

import "math"

//Input: s = "aababbab"
//Output: 2
//Explanation: You can either:
//Delete the characters at 0-indexed positions 2 and 6 ("aababbab" -> "aaabbb"), or
//Delete the characters at 0-indexed positions 3 and 6 ("aababbab" -> "aabbbb").
//func MinimumDeletions(s string) int{
//	var l int = len(s)
//	var a_cnt int = 0
//	for i := 0;i < l;i++{
//		if s[i] == 'a'{
//			a_cnt++
//		}
//	}
//	var b_cnt int = 0
//	var res int = math.MaxInt32
//	for i := 0;i < l;i++{
//		if s[i] == 'b'{
//			res = int(math.Min(float64(res),float64(a_cnt + b_cnt)))
//			b_cnt++
//		}else{
//			a_cnt--
//		}
//	}
//	res = int(math.Min(float64(res),float64(b_cnt)))
//	return res
//}

func MinimumDeletions(s string) int {
	var l int = len(s)
	if l == 1{
		return 0
	}
	var a_cnt []int = make([]int,l)
	var b_cnt []int = make([]int,l)
	var exist_a bool = false
	var exist_b bool = false
	for i := 0;i < l;i++{ //count b from left to right,count a from right to left
		if s[i] == 'b'{
			if i > 0{
				b_cnt[i] = b_cnt[i - 1] + 1
			}else{
				b_cnt[i] = 1
			}
			exist_b = true
		}else{
			if i > 0{
				b_cnt[i] = b_cnt[i  - 1]
			}else{
				b_cnt[i] = 0
			}
			exist_a = true
		}

		if s[l - 1 - i] == 'a'{
			if i != 0{
				a_cnt[l - 1 - i] = a_cnt[l - i] + 1
			}else{
				a_cnt[l - 1 - i] = 1
			}
		}else{
			if i != 0{
				a_cnt[l - 1 - i] = a_cnt[l - i]
			}else{
				a_cnt[l - 1 - i] = 0
			}
		}
	}
	if !exist_a || !exist_b{
		return 0
	}
	var res int = math.MaxInt32
	for i := 0;i < l - 1;i++{
		var min_delete int = 0
		if a_cnt[i + 1] == 0 && b_cnt[i] == 0{
			return min_delete
		}else if a_cnt[i + 1] == 0{
			min_delete = b_cnt[i]
		}else if b_cnt[i] == 0{
			min_delete = a_cnt[i + 1]
		}else{
			//min_delete = int(math.Min(float64(a_cnt[i]), float64(b_cnt[i])))
			if b_cnt[i] == i + 1 && a_cnt[i + 1] == l - i - 1{
				min_delete = int(math.Min(float64(a_cnt[i]), float64(b_cnt[i])))
			}else{
				min_delete = a_cnt[i + 1] + b_cnt[i]
			}
		}
		if min_delete < res{
			res = min_delete
		}
	}
	res = int(math.Min(float64(res), float64(b_cnt[l - 1])))
	res = int(math.Min(float64(res), float64(a_cnt[0])))
	return res
}