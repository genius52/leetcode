package number

import "sort"

func MinimumPushes2(word string) int {
	var cnt [26]int
	for _, w := range word {
		cnt[w-'a']++
	}
	var char_cnt []int
	for i := 0; i < 26; i++ {
		if cnt[i] != 0 {
			char_cnt = append(char_cnt, cnt[i])
		}
	}
	sort.Slice(char_cnt, func(i, j int) bool {
		return char_cnt[i] >= char_cnt[j]
	})
	var res int = 0
	for i := 0; i < len(char_cnt); i++ {
		if i < 8 {
			res += char_cnt[i]
		} else if i >= 8 && i < 16 {
			res += char_cnt[i] * 2
		} else if i >= 16 && i < 24 {
			res += char_cnt[i] * 3
		} else {
			res += char_cnt[i] * 4
		}
	}
	return res
}
