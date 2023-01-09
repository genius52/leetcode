package string_issue

func TakeCharacters(s string, k int) int {
	var l int = len(s)
	var char_cnt [3]int
	for i := 0;i < l;i++{
		char_cnt[s[i] - 'a']++
	}
	if char_cnt[0] < k || char_cnt[1] < k || char_cnt[2] < k{
		return -1
	}
	for i := 0;i < 3;i++{
		char_cnt[i] -= k //允许字符出现的最大次数
	}
	var cur [3]int
	var res int = l
	var left int = 0
	var right int = 0
	for left < l && right < l{
		cur[s[right] - 'a']++
		if cur[s[right] - 'a'] > char_cnt[s[right] - 'a']{
			for left < l && (cur[0] > char_cnt[0] || cur[1] > char_cnt[1] || cur[2] > char_cnt[2]){
				cur[s[left] - 'a']--
				left++
			}
		}
		res = min_int(res,l - (right - left + 1))
		right++
	}
	return res
}