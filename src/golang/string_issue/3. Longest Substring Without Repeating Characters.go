package string_issue

func max_int(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}

func LengthOfLongestSubstring(s string) int{
	var l int = len(s)
	var res int = 0
	var left int = 0
	var right int = 0
	var record map[uint8]int = make(map[uint8]int)
	for left < l{
		for right < l && len(record) == right - left{
			record[s[right]]++
			res = max_int(res,right - left)
			right++
		}
		if right == l {
			if len(record) == right - left{
				res = max_int(res,right - left)
			}
			break
		}
		for left < right && len(record) < (right - left){
			record[s[left]]--
			if record[s[left]] == 0{
				delete(record,s[left])
			}
			left++
		}
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	var start int = 0
	var end int = 0
	var max int = 0
	var cur_len int = 0
	var record map[uint8]int = make(map[uint8]int)
	for end < len(s) && start < len(s){
		if _,ok := record[s[end]];ok{
			start = max_int(record[s[end]] + 1,start)
			record[s[end]] = end
			if cur_len > max{
				max = cur_len
			}
			cur_len = end - start + 1
			end++
		}else{
			cur_len++
			record[s[end]] = end
			end++
		}
	}
	if (end - start) > max{
		max = (end - start)
	}
	return max
}