package string_issue

import "sort"

func MaximumLength2(s string) int {
	var l int = len(s)
	var record [26][]int
	var left int = 0
	var right int = 1
	for left < l {
		for right < l && s[right] == s[left] {
			right++
		}
		cur_len := right - left
		record[int(s[left]-'a')] = append(record[int(s[left]-'a')], cur_len)
		left = right
		right = left + 1
	}
	var res int = -1
	for i := 0; i < 26; i++ {
		var l2 int = len(record[i])
		if l2 == 0 {
			continue
		}
		sort.Ints(record[i])
		//choose the longest substring
		if record[i][l2-1] >= 3 {
			res = max_int(res, record[i][l2-1]-2)
		}
		//choose three longest substring
		if l2 >= 3 {
			res = max_int(res, record[i][l2-3])
		}
		//choose two from first longest substring && choose one from second longest substring
		if l2 >= 2 {
			res = max_int(res, min_int(record[i][l2-1]-1, record[i][l2-2]))
		}
	}
	if res == 0 {
		return -1
	}
	return res
}
