package string_issue

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	var l int = len(word)
	var sub_cnt map[string]int = make(map[string]int)
	for i := 0; i < l; i += k {
		sub_cnt[word[i:i+k]]++
	}
	var res int = l / k
	for _, cnt := range sub_cnt {
		replace_cnt := l/k - cnt
		if replace_cnt < res {
			res = replace_cnt
		}
	}
	return res
}
