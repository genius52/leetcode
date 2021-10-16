package string_issue

func maxPower(s string) int {
	var l int = len(s)
	var left int = 0
	var max_len int = 0
	for left < l{
		var right int = left + 1
		for right < l && s[left] == s[right]{
			right++
		}
		cur := right - left
		if cur > max_len{
			max_len = cur
		}
		left = right
	}
	return max_len
}