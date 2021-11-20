package string_issue

//Input: s = "aababbab"
//Output: 2
//Explanation: You can either:
//Delete the characters at 0-indexed positions 2 and 6 ("aababbab" -> "aaabbb"), or
//Delete the characters at 0-indexed positions 3 and 6 ("aababbab" -> "aabbbb").
//func MinimumDeletions(s string) int {
//	var l int = len(s)
//	if l == 1{
//		return 0
//	}
//	var a_cnt []int = make([]int,l)
//	var b_cnt []int = make([]int,l)
//	var exist_a bool = false
//	var exist_b bool = false
//	for i := 0;i < l;i++{ //count b from left to right,count a from right to left
//		if s[i] == 'b'{
//			if i > 0{
//				b_cnt[i] = b_cnt[i - 1] + 1
//			}else{
//				b_cnt[i] = 1
//			}
//			exist_b = true
//		}else{
//			if i > 0{
//				b_cnt[i] = b_cnt[i  - 1]
//			}else{
//				b_cnt[i] = 0
//			}
//			exist_a = true
//		}
//
//		if s[l - 1 - i] == 'a'{
//			if i != 0{
//				a_cnt[l - 1 - i] = a_cnt[l - i] + 1
//			}else{
//				a_cnt[l - 1 - i] = 1
//			}
//		}else{
//			if i != 0{
//				a_cnt[l - 1 - i] = a_cnt[l - i]
//			}else{
//				a_cnt[l - 1 - i] = 0
//			}
//		}
//	}
//	if !exist_a || !exist_b{
//		return 0
//	}
//	var res int = math.MaxInt32
//	for i := 0;i < l - 1;i++{
//		var min_delete int = 0
//		if a_cnt[i + 1] == 0 && b_cnt[i] == 0{
//			return min_delete
//		}else if a_cnt[i + 1] == 0{
//			min_delete = b_cnt[i]
//		}else if b_cnt[i] == 0{
//			min_delete = a_cnt[i + 1]
//		}else{
//			//min_delete = int(math.Min(float64(a_cnt[i]), float64(b_cnt[i])))
//			if b_cnt[i] == i + 1 && a_cnt[i + 1] == l - i - 1{
//				min_delete = int(math.Min(float64(a_cnt[i]), float64(b_cnt[i])))
//			}else{
//				min_delete = a_cnt[i + 1] + b_cnt[i]
//			}
//		}
//		if min_delete < res{
//			res = min_delete
//		}
//	}
//	res = int(math.Min(float64(res), float64(b_cnt[l - 1])))
//	res = int(math.Min(float64(res), float64(a_cnt[0])))
//	return res
//}

//Input: s = "aababbab"
//Output: 2
func MinimumDeletions(s string) int{
	var l int = len(s)
	var steps int = 0
	var b_cnt int = 0
	for i := 0;i < l;i++{
		if s[i] == 'a'{
			//1.删除左边所有的a,包括当前的a
			//1.删除左边所有的b
			steps = min_int(steps + 1,b_cnt)
		}else{
			b_cnt++
		}
	}
	return steps
}

func minimumDeletions(s string) int{
	var l int = len(s)
	var prefix_a []int = make([]int,l + 1)
	var suffix_b []int = make([]int,l)
	for i := 0;i < l;i++{
		if s[i] == 'a'{
			prefix_a[i + 1] = prefix_a[i] + 1
		}else{
			prefix_a[i + 1] = prefix_a[i]
		}
	}
	if s[l - 1] == 'b'{
		suffix_b[l - 1] = 1
	}
	for i := l - 2;i >= 0;i--{
		if s[i] == 'b'{
			suffix_b[i] = suffix_b[i + 1] + 1
		}else{
			suffix_b[i] = suffix_b[i + 1]
		}
	}
	var max_len int = 0
	for i := 0;i < l;i++{
		max_len = max_int(max_len,prefix_a[i + 1] + suffix_b[i])
	}
	return l - max_len
}