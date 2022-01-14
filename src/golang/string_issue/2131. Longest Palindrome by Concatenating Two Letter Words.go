package string_issue

import "sort"

//Input: words = ["ab","ty","yt","lc","cl","ab"]
//Output: 8
//Explanation: One longest palindrome is "ty" + "lc" + "cl" + "yt" = "tylcclyt", of length 8.
//Note that "lcyttycl" is another longest palindrome that can be created.
func LongestPalindrome2131(words []string) int {
	var l int = len(words)
	var record map[string]int = make(map[string]int)
	for i := 0; i < l; i++ {
		record[words[i]]++
	}
	var same_cnt []int
	var palin_cnt int = 0
	for word, cnt1 := range record {
		if word[0] == word[1] {
			same_cnt = append(same_cnt, cnt1)
		} else {
			other := string(word[1]) + string(word[0])
			if cnt2, ok := record[other]; ok {
				palin_cnt += min_int(cnt1, cnt2)
				delete(record, other)
			}
		}
	}
	sort.Ints(same_cnt)
	var res int = palin_cnt * 4
	var same_cnt_len int = len(same_cnt)
	if same_cnt_len == 1 {
		res += same_cnt[0] * 2
	} else if same_cnt_len > 1 {
		var exist_odd bool = false
		for i := same_cnt_len - 1; i >= 0; i-- {
			if (same_cnt[i] | 1) == same_cnt[i] {
				exist_odd = true
			}
			res += (same_cnt[i] / 2 * 2) * 2
		}
		if exist_odd {
			res += 2
		}
	}
	return res
}
