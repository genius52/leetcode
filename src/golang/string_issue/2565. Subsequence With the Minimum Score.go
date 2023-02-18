package string_issue

func MinimumScore(s string, t string) int {
	var l1 int = len(s)
	var l2 int = len(t)
	//left[i]为s[:i]对应的t的最长前缀的结束下标
	//right[i]为s[i:]对应的t的最长后缀的开始下标
	//删除的子串就是从 left[i]+1 到 right[i]−1 这段，答案就是 right[i]−left[i]−1 的最小值
	var left []int = make([]int, l1)
	var right []int = make([]int, l1)
	var j int = 0
	for i := 0; i < l1; i++ {
		if j < l2 && s[i] == t[j] {
			j++
		}
		left[i] = j
	}
	j = l2 - 1
	for i := l1 - 1; i >= 0; i-- {
		if j >= 0 && s[i] == t[j] {
			j--
		}
		right[i] = j
	}
	var res int = 1<<31 - 1
	for i := 0; i < l1-1; i++ {
		left_len := left[i]
		right_len := l2 - right[i+1] - 1
		if left_len+right_len >= l2 {
			return 0
		}
		res = min_int(res, l2-left_len-right_len)
	}
	res = min_int(res, l2-left[l1-1])
	res = min_int(res, right[0])
	return res
}
