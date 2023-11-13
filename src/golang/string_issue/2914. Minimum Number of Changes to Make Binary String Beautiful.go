package string_issue

func MinChanges(s string) int {
	var l int = len(s)
	var res int = 0
	var one_cnt int = 0
	for i := 0; i < l; i++ {
		if s[i] == '1' {
			one_cnt++
		}
		if i|1 == i { //idx == 1,idx == 3...
			if one_cnt == 1 {
				res++
			}
			one_cnt = 0
		}
	}
	return res
}

//func MinChanges(s string) int {
//	var l int = len(s)
//	var prefix_one_cnt []int = make([]int, l)
//	var suffix_one_cnt []int = make([]int, l)
//	var one_cnt int = 0
//	for i := 0; i < l; i++ {
//		if s[i] == '1' {
//			one_cnt++
//		}
//		prefix_one_cnt[i] = one_cnt
//	}
//	one_cnt = 0
//	for i := l - 1; i >= 0; i-- {
//		if s[i] == '1' {
//			one_cnt++
//		}
//		suffix_one_cnt[i] = one_cnt
//	}
//	var res int = l
//	for i := 0; i < l; i++ {
//		if i == 0 {
//			res = min_int(res, suffix_one_cnt[i])   //后缀全改成0
//			res = min_int(res, l-suffix_one_cnt[i]) //后缀全改成1
//		} else if i == l-1 {
//			res = min_int(res, prefix_one_cnt[i])     //前缀全改成0
//			res = min_int(res, i+1-prefix_one_cnt[i]) //前缀全改成1
//		} else {
//			res = min_int(res, prefix_one_cnt[i-1]+suffix_one_cnt[i])       //前缀和后缀全改成0
//			res = min_int(res, i+1-prefix_one_cnt[i-1]+l-suffix_one_cnt[i]) //前缀和后缀全改成1
//		}
//	}
//	return res
//}
