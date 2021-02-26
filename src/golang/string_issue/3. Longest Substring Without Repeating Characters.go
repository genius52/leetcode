package string_issue

func max_int(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
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