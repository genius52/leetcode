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