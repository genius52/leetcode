package string_issue

func maxLengthBetweenEqualCharacters(s string) int {
	var cnt [26]int
	for i := 0;i < 26;i++{
		cnt[i] = -1
	}
	var res int = -1
	for idx,c := range s {
		if cnt[c - 'a'] == -1{
			cnt[c - 'a'] = idx
		}else{
			dis := idx - cnt[c - 'a'] - 1
			if dis > res{
				res = dis
			}
		}
	}
	return res
}