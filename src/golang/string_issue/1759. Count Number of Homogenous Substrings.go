package string_issue

//Input: s = "abbcccaa"
//Output: 13
//Explanation: The homogenous substrings are listed as below:
//"a"   appears 3 times.
//"aa"  appears 1 time.
//"b"   appears 2 times.
//"bb"  appears 1 time.
//"c"   appears 3 times.
//"cc"  appears 2 times.
//"ccc" appears 1 time.
//3 + 1 + 2 + 1 + 3 + 2 + 1 = 13.
func countHomogenous(s string) int {
	var l int = len(s)
	var res int = 0
	var cnt int = 1
	for i := 0; i < l; i++ {
		if i > 0 {
			if s[i] == s[i-1] {
				cnt++
			} else {
				cnt = 1
			}
		}
		res += cnt
		res %= 1e9 + 7
	}
	return res
}

func calculate(n int) int {
	return n + n*(n-1)/2
}

func CountHomogenous(s string) int {
	var l int = len(s)
	if l <= 1 {
		return 1
	}
	s += "0"
	var begin int = 0
	var end int = 1
	var res int = 0
	for end <= l {
		if s[end] == s[end-1] {
			end++
		} else {
			res += calculate(end - begin)
			res %= 1000000007
			begin = end
			end++
		}
	}
	return res
}
