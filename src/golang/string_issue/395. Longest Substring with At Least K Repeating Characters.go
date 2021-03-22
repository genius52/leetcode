package string_issue

func LongestSubstring(s string, k int) int{
	l := len(s)
	var less_than_k map[uint8]int = make(map[uint8]int)
	for i := 0;i < l;i++{
		if _,ok := less_than_k[s[i]];ok{
			less_than_k[s[i]]++
		}else{
			less_than_k[s[i]] = 1
		}
	}
	for c,n := range less_than_k{
		if n >= k{
			delete(less_than_k,c)
		}
	}
	if len(less_than_k) == 0{
		return l
	}
	if len(less_than_k) == l{
		return 0
	}
	var trace []int = make([]int,l)
	for i := 0;i < l;i++{
		if _,ok := less_than_k[s[i]];ok{
			trace[i] = -1
		}
	}
	var max_len = 0
	start := 0
	for start < l {
		for start < l{
			if trace[start] != - 1{
				break
			}
			start++
		}
		if start == l{
			break
		}
		end := start
		for end < l{
			if trace[end] == -1{
				break
			}
			end++
		}
		res := LongestSubstring(s[start:end],k)
		if res > max_len{
			max_len = res
		}
		start++
	}
	return max_len
}