package string_issue

//Input: s = "aabcb"
//Output: 5
//Explanation: The substrings with non-zero beauty are ["aab","aabc","aabcb","abcb","bcb"], each with beauty equal to 1.
func calcu_beauty(freq [26]int)int{
	var low int = 2147483647
	var high int = 0
	for i := 0;i < 26;i++{
		if freq[i] == 0{
			continue
		}
		if freq[i] > high{
			high = freq[i]
		}
		if freq[i] < low{
			low = freq[i]
		}
	}
	return high - low
}

func BeautySum(s string) int {
	var l int = len(s)
	var res int = 0
	for i := 0;i < l;i++{
		var frequncy [26]int
		frequncy[s[i] - 'a']++
		for j := i + 1;j < l;j++{
			frequncy[s[j] - 'a']++
			res += calcu_beauty(frequncy)
		}
	}
	return res
}