package string_issue

func rearrangeCharacters(s string, target string) int {
	var record [26]int
	for _,c := range s{
		record[c - 'a']++
	}
	var t map[int]int = make(map[int]int)
	for _,c := range target{
		t[int(c - 'a')]++
	}
	var res int = 2147483647
	for c,cnt1 := range t{
		cnt2 := record[c]
		n := cnt2 / cnt1
		if n < res{
			res = n
		}
	}
	if res == 2147483647{
		return 0
	}
	return res
}