package string_issue

func FirstUniqChar(s string) int{
	var cnt [26]int
	for i := 0;i < 26;i++{
		cnt[i] = -1
	}
	for idx,c := range s{
		if cnt[c - 'a'] != -1{
			cnt[c - 'a'] = -2
		}else{
			cnt[c - 'a'] = idx
		}
	}
	var res int = 2147483647
	for i := 0;i < 26;i++{
		if cnt[i] != -1 && cnt[i] != -2{
			if cnt[i] < res{
				res = cnt[i]
			}
		}
	}
	if res == 2147483647{
		return -1
	}
	return res
}

//func firstUniqChar(s string) int {
//	var record [26]int
//	var l int = len(s)
//	for i := 0;i < l;i++{
//		record[s[i] - 'a']++
//	}
//	var res int = -1
//	for i := 0;i < l;i++{
//		if record[s[i] - 'a'] == 1{
//			res = i
//			break
//		}
//	}
//	return res
//}