package string_issue

//Input:
//s: "cbaebabacd" p: "abc"
//
//Output:
//[0, 6]
func findAnagrams(s string, p string) []int {
	var p_len int = len(p)
	var p_cnt [26]int
	for i := 0;i < p_len;i++{
		p_cnt[p[i] - 'a']++
	}
	var res []int
	var l int = len(s)
	if l < p_len{
		return res
	}
	var cur_cnt [26]int
	var left int = 0
	var right int = p_len - 1
	for i := left;i <= right;i++{
		cur_cnt[s[i] - 'a']++
	}
	for left < l && right < l{
		var same bool = true
		for i := 0;i < 26;i++{
			if cur_cnt[i] != p_cnt[i]{
				same = false
				break
			}
		}
		if same{
			res = append(res,left)
		}
		cur_cnt[s[left] - 'a']--
		left++
		right++
		if right < l{
			cur_cnt[s[right] - 'a']++
		}
	}
	return res
}