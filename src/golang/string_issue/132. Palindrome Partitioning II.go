package string_issue

//Input: s = "aab"
//Output: 1
//Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.
func dp_minCut(s string,start int,end int,memo map[int]int)int{
	if start >= end{
		return 0
	}
	if cnt,ok := memo[start];ok{
		return cnt
	}
	var min_len int = 1 << 31
	for i := start + 1;i <= end;i++{
		var cur_len int = 0
		if isPalindrome(s[start:i]){
			if i == end{
				cur_len = 0
			}else{
				cur_len = 1 + dp_minCut(s,i,end,memo)
			}
			if cur_len < min_len{
				min_len = cur_len
			}
		}
	}
	memo[start] = min_len
	return min_len
}

func MinCut(s string) int {
	var l int = len(s)
	if l <= 1{
		return 0
	}
	var memo map[int]int = make(map[int]int)
	return dp_minCut(s,0,l,memo)
}