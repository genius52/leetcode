package string_issue

// S = "cba"
// T = "abcd"
// Output: "cbad"
func CustomSortString(S string, T string) string {
	var records [26]uint8
	var l1 int = len(S)
	for i := 0;i < l1;i++{
		records[i] = S[i]//1st : c, 2nd : b, 3rd: a
	}
	var l2 int = len(T)
	var char_cnt [26]int
	for i := 0;i < l2;i++{
		char_cnt[T[i] - 'a']++
	}
	var res string
	for i := 0;i < 26;i++{
		if records[i] == 0{
			continue
		}
		if char_cnt[records[i] - 'a'] == 0{
			continue
		}
		cnt := char_cnt[records[i] - 'a']
		char_cnt[records[i] - 'a'] = 0
		for j := 0;j < cnt;j++{
			res += string(records[i])
		}
	}
	for i := 0;i < 26;i++{
		if char_cnt[i] == 0{
			continue
		}
		cnt := char_cnt[i]
		for j := 0;j < cnt;j++{
			c := i + 'a'
			res += string(c)
		}
	}
	return res
}
