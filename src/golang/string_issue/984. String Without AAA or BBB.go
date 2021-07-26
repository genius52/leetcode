package string_issue

//Input: a = 1, b = 2
//Output: "abb"
//Explanation: "abb", "bab" and "bba" are all correct answers.
func strWithout3a3b(a int, b int) string {
	var res string
	var s1 string
	var s2 string
	var diff int = 0
	var min_cnt int = 0
	if a > b {
		min_cnt = b
		diff = a - b
		s1 = "a"
		s2 = "b"
	}else if b > a{
		min_cnt = a
		diff = b - a
		s1 = "b"
		s2 = "a"
	}else{
		min_cnt = a
		s1 = "a"
		s2 = "b"
	}
	for diff > 0 && min_cnt > 0{
		res += s1 + s1 + s2
		diff--
		min_cnt--
	}
	for diff > 0{
		res += s1
		diff--
	}
	for min_cnt > 0{
		res += s1 + s2
		min_cnt--
	}
	return res
}