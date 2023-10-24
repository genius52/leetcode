package string_issue

func dp_minOperations2896(s1 string, s2 string, l int, idx int, reverse_cnt int, pre_reverse bool, x int, memo [][][2]int) int {
	if idx == l {
		if pre_reverse || reverse_cnt > 0 {
			return -1
		}
		return 0
	}
	if pre_reverse {
		if memo[idx][reverse_cnt][1] != -2 {
			return memo[idx][reverse_cnt][1]
		}
	} else {
		if memo[idx][reverse_cnt][0] != -2 {
			return memo[idx][reverse_cnt][0]
		}
	}
	var res int = 1<<31 - 1
	if ((s1[idx] == s2[idx]) && !pre_reverse) || ((s1[idx] != s2[idx]) && pre_reverse) {
		//不需要反转
		ret0 := dp_minOperations2896(s1, s2, l, idx+1, reverse_cnt, false, x, memo)
		if ret0 != -1 {
			res = min_int(res, ret0)
		}
	}
	//反转当前和下一个
	ret1 := dp_minOperations2896(s1, s2, l, idx+1, reverse_cnt, true, x, memo)
	if ret1 != -1 {
		res = min_int(res, 1+ret1)
	}
	//反转当前和下下个其他的位置
	ret2 := dp_minOperations2896(s1, s2, l, idx+1, reverse_cnt+1, false, x, memo)
	if ret2 != -1 {
		res = min_int(res, x+ret2)
	}
	if reverse_cnt > 0 {
		//免费反转
		ret3 := dp_minOperations2896(s1, s2, l, idx+1, reverse_cnt-1, false, x, memo)
		if ret3 != -1 {
			res = min_int(res, ret3)
		}
	}
	if res == (1<<31 - 1) {
		if pre_reverse {
			memo[idx][reverse_cnt][1] = -1
		} else {
			memo[idx][reverse_cnt][0] = -1
		}
		return -1
	} else {
		if pre_reverse {
			memo[idx][reverse_cnt][1] = res
		} else {
			memo[idx][reverse_cnt][0] = res
		}
		return res
	}
}

func MinOperations2896(s1 string, s2 string, x int) int {
	var l int = len(s1)
	var memo [][][2]int = make([][][2]int, l)
	for i := 0; i < l; i++ {
		memo[i] = make([][2]int, l+1)
		for j := 0; j <= l; j++ {
			memo[i][j][0] = -2
			memo[i][j][1] = -2
		}
	}
	var diff int
	for i := 0; i < l; i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	if diff%2 == 1 {
		return -1
	}
	return dp_minOperations2896(s1, s2, l, 0, 0, false, x, memo)
}

//func minOperations2896(s1 string, s2 string, x int) int {
//	var l int = len(s1)
//	var reverse_cnt int = 0
//	var cur_reverse bool = false
//	var res int = 0
//	for i := 0;i < l;i++{
//		if s1[i] == s2[i]{
//			if cur_reverse{
//
//			}else{
//
//			}
//		}else{
//
//		}
//	}
//	return res
//}
