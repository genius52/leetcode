package string_issue

import "strconv"

func countPalindromicSubsequence(s string) int {
	var l int = len(s)
	var record map[string]bool = make(map[string]bool)
	var char_idx [26][]int
	for i := 0;i < l;i++{
		cur_char := s[i] - 'a'
		if len(char_idx[cur_char]) > 0{
			pre_idx := char_idx[cur_char][0]
			for j := 0;j < 26;j++{
				var l2 int = len(char_idx[j])
				if l2 == 0{
					continue
				}
				if char_idx[j][l2 - 1] <= pre_idx{
					continue
				}
				//record[int(cur_char + 1) + (j + 1) * 26 +  int((cur_char + 1) * 26 * 26)] = true
				record[strconv.Itoa(int(cur_char)) + strconv.Itoa(j) + strconv.Itoa(int(cur_char))] = true
			}
		}
		char_idx[cur_char] = append(char_idx[cur_char],i)
	}
	return len(record)
}