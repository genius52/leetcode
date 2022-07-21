package string_issue

import "math"

//76
//Input: S = "ADOBECODEBANC", T = "ABC"
//Output: "BANC"
func MinWindow(s string, t string) string {
	var record map[uint8]int = make(map[uint8]int)
	l := len(s)
	for i := 0;i < len(t);i++{
		if _,ok := record[t[i]];ok{
			record[t[i]]++
		}else{
			record[t[i]] = 1
		}
	}
	var counter int = len(t)
	begin := 0
	end := 0
	var res string
	var min_len int = math.MaxInt32
	for begin < l && end < l{
		if _,ok := record[s[end]];ok{
			if record[s[end]] > 0{
				counter--
			}
			record[s[end]]--
		}
		end++
		for counter == 0 {
			if (end - begin) < min_len{
				min_len = end - begin
				res = s[begin:end]
			}
			if _,ok := record[s[begin]];ok{
				record[s[begin]]++
				if record[s[begin]] > 0{
					counter++
				}
			}
			begin++
		}
	}
	return res
}

func minWindow(s string, t string) string{
	var total int = len(t)
	var target map[uint8]int = make(map[uint8]int)
	for i := 0;i < total;i++{
		target[t[i]]++
	}
	var left int = 0
	var right int = 0
	var l int = len(s)
	var res string
	var min_len int = 2147483647
	for left < l && right < l{
		for right < l && total > 0{
			if cnt,ok := target[s[right]];ok{
				if cnt > 0{
					total--
				}
				target[s[right]]--
			}
			right++
		}
		for total == 0{
			cur_len := right - left + 1
			if cur_len < min_len{
				min_len = cur_len
				res = s[left:right]
			}
			if _,ok := target[s[left]];ok{
				target[s[left]]++
				if target[s[left]] > 0{
					total++
				}
			}
			left++
		}
	}
	return res
}