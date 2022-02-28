package string_issue

func minSteps2186(s string, t string) int {
	var cnt1 [26]int
	var cnt2 [26]int
	for _,c := range s{
		cnt1[c - 'a']++
	}
	for _,c := range t{
		cnt2[c - 'a']++
	}
	var res int = 0
	for i := 0;i < 26;i++{
		res += abs_int(cnt1[i] - cnt2[i])
	}
	return res
}