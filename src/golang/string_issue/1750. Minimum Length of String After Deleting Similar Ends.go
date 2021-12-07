package string_issue

//Input: s = "aabccabba"
//Output: 3
//Explanation: An optimal sequence of operations is:
//- Take prefix = "aa" and suffix = "a" and remove them, s = "bccabb".
//- Take prefix = "b" and suffix = "bb" and remove them, s = "cca".
func MinimumLength(s string) int {
	var l int = len(s)
	var left int = 0
	var right int = l - 1
	for left < right && s[left] == s[right] {
		var i int = left + 1
		for i < right && s[i] == s[left] {
			i++
		}
		left = i
		var j int = right - 1
		for j > left && s[j] == s[right] {
			j--
		}
		right = j
	}
	return right - left + 1
}

//func MinimumLength(s string) int {
//	var l int = len(s)
//	var begin int = 0
//	var end int = l - 1
//	for begin < end {
//		if s[begin] != s[end] {
//			break
//		}
//		var begin_move int = begin
//		for begin_move < end && s[begin_move] == s[begin] {
//			begin_move++
//		}
//		if begin_move == end && s[begin_move] == s[end] {
//			return 0
//		}
//		var end_move int = end
//		for end_move > begin && s[end_move] == s[end] {
//			end_move--
//		}
//		if end_move == begin && s[end_move] == s[begin] {
//			return 0
//		}
//		begin = begin_move
//		end = end_move
//	}
//	return end - begin + 1
//}
