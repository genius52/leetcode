package string_issue

func divideString(s string, k int, fill byte) []string {
	var l int = len(s)
	var res []string
	var left int = 0
	var right int = k
	for left < l{
		res = append(res,s[left:min_int(right,l)])
		left += k
		right += k
	}
	if l % k != 0{
		var l2 int = len(res[len(res) - 1])
		for l2 < k{
			res[len(res) - 1] += string(fill)
			l2++
		}
	}
	return res
}