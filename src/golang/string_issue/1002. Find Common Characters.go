package string_issue

//Input: ["bella","label","roller"]
//Output: ["e","l","l"]
func CommonChars(words []string) []string {
	var l int = len(words)
	var cnt [][26]int = make([][26]int,l)
	var res []string
	for i := 0;i < l;i++{
		for _,c := range words[i]{
			cnt[i][c - 'a']++
		}
	}
	for i := 0;i < 26;i++{
		var min_cnt int = 2147483647
		for j := 0;j < l;j++{
			if cnt[j][i] == 0{
				min_cnt = 0
				break
			}else{
				if cnt[j][i] < min_cnt{
					min_cnt = cnt[j][i]
				}
			}
		}
		for min_cnt > 0{
			s := string(i + 'a')
			res = append(res,s)
			min_cnt--
		}
	}
	return res
}