package string_issue

import "math"

//Input: "00011000"
//Output: 2
//Explanation: We flip to get 00000000.
func MinFlipsMonoIncr(S string) int{
	S = "0" + S
	var l int = len(S)
	if l == 2{
		return 0
	}
	visit := 0
	for visit < l{
		if S[visit] == '0'{
			visit++
		}else{
			break
		}
	}
	if visit == l || visit == l - 1{
		return 0
	}
	var left_one_cnt []int = make([]int,l)
	var right_one_cnt []int = make([]int,l)
	if S[visit] == '1'{
		left_one_cnt[visit] = 1
	}
	if S[l - 1] == '1'{
		right_one_cnt[l - 1] = 1
	}
	for i := visit + 1;i < l;i++ {
		if S[i] == '1' {
			left_one_cnt[i] = left_one_cnt[i-1] + 1
		} else {
			left_one_cnt[i] = left_one_cnt[i-1]
		}
	}
	for i := 1;i < l;i++{
		if S[l-1-i] == '1' {
			right_one_cnt[l-1-i] = right_one_cnt[l-i] + 1
		} else {
			right_one_cnt[l-1-i] = right_one_cnt[l-i]
		}
	}

	var res int = math.MaxInt32
	for i := visit + 1;i < l;i++{
		//1: 1 -> 0 on left side,0 -> 1 on right side
		//2: 0 -> 1 on left side,0 -> 1 on right side
		//3: 1 -> 0 on left side,1 -> 0 on right side
		var left_zero_cnt int = 0
		var right_zero_cnt int = 0
		if i == l - 1{
			left_zero_cnt = i  + 1 - visit - left_one_cnt[i]
			right_zero_cnt = 0
		}else{
			left_zero_cnt = i  + 1 - visit - left_one_cnt[i]
			right_zero_cnt = l - i - 1 - right_one_cnt[i + 1]
		}
		cur_steps := int(math.Min(float64(left_one_cnt[i] + right_zero_cnt),float64(left_zero_cnt + right_zero_cnt)))
		if i == l - 1{
			cur_steps = int(math.Min(float64(cur_steps),float64(left_one_cnt[i])))
		}else{
			cur_steps = int(math.Min(float64(cur_steps),float64(left_one_cnt[i] + right_one_cnt[i + 1])))
		}

		if cur_steps < res{
			res = cur_steps
		}
	}
	return res
}

//func dfs_minFlipsMonoIncr(S string,pos int,pre_exist_one bool) int{
//	if pos >= len(S){
//		return 0
//	}
//	if pre_exist_one{
//		if S[pos] == '0'{
//			return 1 + dfs_minFlipsMonoIncr(S,pos + 1,true)
//		}else{
//			return  dfs_minFlipsMonoIncr(S,pos + 1,true)
//		}
//	}else{
//		if S[pos] == '0'{
//			return  dfs_minFlipsMonoIncr(S,pos + 1,false)
//		}else{
//			res1 := 1 + dfs_minFlipsMonoIncr(S,pos + 1,false)//make 1 -> 0
//			res2 := dfs_minFlipsMonoIncr(S,pos + 1,true)
//			return int(math.Min(float64(res1),float64(res2)))
//		}
//	}
//}
//
//func minFlipsMonoIncr(S string) int {
//	S = "0" + S
//	var l int = len(S)
//	visit := 0
//	for visit < l{
//		if S[visit] == '0'{
//			visit++
//		}else{
//			break
//		}
//	}
//	if visit < l{
//		return dfs_minFlipsMonoIncr(S,visit - 1,false)
//	}else{
//		return 0
//	}
//}