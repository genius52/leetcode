package string_issue

func areOccurrencesEqual(s string) bool {
	var record [26]int
	for _,c := range s{
		record[c - 'a']++
	}
	var pre_cnt int = 0
	for i := 0;i < 26;i++{
		if record[i] == 0{
			continue
		}
		if pre_cnt == 0{
			pre_cnt = record[i]
		}else{
			if pre_cnt != record[i]{
				return false
			}
		}
	}
	return true
}