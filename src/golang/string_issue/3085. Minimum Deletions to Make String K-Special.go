package string_issue

import "sort"

func MinimumDeletions3085(word string, k int) int {
	var char_cnt [26]int
	for _, c := range word {
		char_cnt[c-'a']++
	}
	var record []int
	for i := 0; i < 26; i++ {
		if char_cnt[i] > 0 {
			record = append(record, char_cnt[i])
		}
	}
	sort.Ints(record)
	var total_cnt int = 0
	for _, n := range record {
		total_cnt += n
	}
	//所有数字的范围都在low和high之间
	var low int = record[0]
	var high int = low + k
	var l int = len(record)
	if l == 1 || high >= record[l-1] {
		return 0
	}
	var res int = total_cnt
	for high <= record[l-1] {
		//find_idx := sort.Search(l, func(i int) bool { //最小的索引，使得record[i] > high成立
		//	return record[i] > high
		//})
		var delete_cnt int = 0
		for i := 0; i < l; i++ {
			if record[i] < low {
				delete_cnt += record[i]
			} else if record[i] > high {
				delete_cnt += record[i] - high
			}
		}
		if delete_cnt < res {
			res = delete_cnt
		}
		low++
		high++
	}
	return res
}
