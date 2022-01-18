package string_issue

import "strconv"

func CountPalindromicSubsequence(s string) int{
	var l int = len(s)
	var pre [][26]int = make([][26]int,l)
	pre[0][s[0] - 'a']++
	var record map[int]bool = make(map[int]bool)
	for i := 1;i < l;i++{
		for j := 0;j < 26;j++{
			pre[i][j] = pre[i - 1][j]
		}
		pre[i][s[i] - 'a']++
	}
	var res int = 0
	for i := 1;i < l - 1;i++{
		for j := 0;j < 26;j++{
			left_cnt := pre[i - 1][j]
			if left_cnt == 0{
				continue
			}
			right_cnt := pre[l - 1][j] - pre[i][j]
			if right_cnt == 0{
				continue
			}
			key := j * 26 + int(s[i] - 'a')
			if _,ok := record[key];!ok{
				record[key] = true
				res++
			}
		}
	}
	return res
}

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