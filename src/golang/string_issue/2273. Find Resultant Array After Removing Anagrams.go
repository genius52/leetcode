package string_issue

func removeAnagrams(words []string) []string {
	var l int = len(words)
	var res []string
	var left int = 0
	for left < l{
		res = append(res,words[left])
		var l1 int = len(words[left])
		var cnt1 [26]int
		for i := 0;i < l1;i++{
			cnt1[words[left][i] - 'a']++
		}
		var right int = left + 1
		for right < l{
			var l2 int = len(words[right])
			if l1 != l2{
				break
			}
			var cnt2 [26]int
			for j := 0;j < l2;j++{
				cnt2[words[right][j] - 'a']++
			}
			var equal bool = true
			for i := 0;i < 26;i++{
				if cnt1[i] != cnt2[i]{
					equal = false
					break
				}
			}
			if !equal{
				break
			}
			right++
		}
		left = right
	}
	return res
}