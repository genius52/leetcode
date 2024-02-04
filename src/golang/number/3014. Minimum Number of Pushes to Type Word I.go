package number

import "sort"

func minimumPushes(word string) int {
	var char_cnt map[int32]int = make(map[int32]int)
	for _, c := range word {
		char_cnt[c]++
	}
	var freq_cnt []int
	for _, cnt := range char_cnt {
		freq_cnt = append(freq_cnt, cnt)
	}
	sort.Slice(freq_cnt, func(i, j int) bool {
		return i >= j
	})
	var res int = 0
	for i := 0; i < len(freq_cnt); i++ {
		if i < 8 {
			res += freq_cnt[i]
		} else if i >= 8 && i < 16 {
			res += freq_cnt[i] * 2
		} else if i >= 16 && i < 24 {
			res += freq_cnt[i] * 3
		} else {
			res += freq_cnt[i] * 4
		}
	}
	return res
}
