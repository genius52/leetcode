package string_issue

//Input: s = "3242415"
//Output: 5
//Explanation: "24241" is the longest awesome substring, we can form the palindrome "24142" with some swaps.
func LongestAwesome(s string) int {
	var l int = len(s)
	var record map[int]int = make(map[int]int)
	record[0] = -1
	var res int = 0
	var mask int = 0
	for i := 0;i < l;i++{
		offset := s[i] - '0'
		mask = mask ^ (1 << offset) //原来是奇数，现在变偶数;原来是偶数，现在变奇数
		for k,v := range record{
			n := mask ^ k
			if n == 0 || n & (n - 1) == 0{//0-9，奇数和偶数的个数都相同;0-9，奇数和偶数的个数只有一个不同
				res = max_int(res,i - v)
			}
		}
		if _,ok := record[mask];!ok{
			record[mask] = i
		}
	}
	return res
}