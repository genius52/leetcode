package string_issue

import "math"

//Input: s = "111000"
//Output: 1
//Explanation: Swap positions 1 and 4: "111000" -> "101010"
//The string is now alternating.
func calc_wrongpos(s string,first uint8)int{
	var l int = len(s)
	var res int = 0
	for i := 0;i < l;i += 2{
		if s[i] != first{
			res++
		}
	}
	return res
}

func minSwaps(s string) int {
	var l int = len(s)
	var one_cnt int = 0
	var zero_cnt int = 0
	for i := 0;i < l;i++{
		if s[i] == '0'{
			zero_cnt++
		}else{
			one_cnt++
		}
	}
	var res int = 2147483647
	if (l | 1) == l{
		if math.Abs(float64(zero_cnt - one_cnt)) != 1{
			return -1
		}
		if zero_cnt > one_cnt{
			res = calc_wrongpos(s,'0')
		}else{
			res = calc_wrongpos(s,'1')
		}
	}else{
		if zero_cnt != one_cnt{
			return -1
		}
		zero_one := calc_wrongpos(s,'0')
		one_zero := calc_wrongpos(s,'1')
		if zero_one < one_zero{
			res = zero_one
		}else{
			res = one_zero
		}
	}
	return res
}