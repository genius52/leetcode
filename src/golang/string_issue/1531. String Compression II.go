package string_issue

import "strconv"

func calc_len(s string)int{
	var res int = 0
	var left int = 0
	var l int = len(s)
	for left < l{
		 var right int = left + 1
		 for right < l && s[right] == s[left]{
		 	right++
		 }
		 sub_len := right - left
		 if sub_len > 1{
		 	for sub_len > 0{
		 		res++
		 		sub_len /= 10
			}
		 }
		 res++
		 left = right
	}
	return res
}

func dp_getLengthOfOptimalCompression(s string,pos int,k int,memo map[string]int)int{
	if k == 0{
		return calc_len(s)
	}
	key := s + strconv.Itoa(k)
	if _,ok := memo[key];ok{
		return memo[key]
	}
	var res int = 2147483647
	var left int = pos
	var l int = len(s)
	for left < l{
		var right int = left + 1
		for right < l && s[right] == s[left]{
			right++
		}
		sub_len := right - left
		if sub_len <= k{
			s2 := s[:left] + s[right:]
			res = min_int(res,dp_getLengthOfOptimalCompression(s2,left,k - sub_len,memo))
		}else{
			s2 := s[:left] + s[left + k:]
			var cur_res int = calc_len(s2)
			cur_key := s2 + strconv.Itoa(0)
			memo[cur_key] = cur_res
			res = min_int(res,cur_res)
		}
		left = right
	}
	memo[key] = res
	return res
}

func GetLengthOfOptimalCompression(s string, k int) int{
	var memo map[string]int = make(map[string]int)
	return dp_getLengthOfOptimalCompression(s,0,k,memo)
}