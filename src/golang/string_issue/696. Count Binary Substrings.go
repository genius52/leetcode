package string_issue

//Input: "00110011"
//Output: 6
//Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
func is_binary(s string,start,end int,l int)int{
	if start < 0 || end >= l{
		return 0
	}
	if s[start] == s[end]{
		return 0
	}
	var cnt int = 1
	head_val := s[start]
	tail_val := s[end]
	start--
	end++
	for start >= 0 && end < l{
		if s[start] != head_val || s[end] != tail_val{
			return cnt
		}
		start--
		end++
		cnt++
	}
	return cnt
}

func CountBinarySubstrings(s string) int {
	var res int = 0
	l := len(s)
	for i := 0;i < (l - 1);i++{
		res += is_binary(s,i,i+1,l)
	}
	return res
}