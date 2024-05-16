package string_issue

func dp_minimumSubstringsInPartition(s string, idx int, l int, char_cnt []int, memo []int) int {
	if idx == len(s) {
		return 0
	}
	if memo[idx] != -1 {
		return memo[idx]
	}
	var res int = 1<<31 - 1
	var max_cnt int = 0
	var letter_cnt int = 0
	for i := idx; i < l; i++ {
		char_cnt[s[i]-'a']++
		if char_cnt[s[i]-'a'] == 1 {
			letter_cnt++
		}
		max_cnt = max_int(max_cnt, char_cnt[s[i]-'a'])
		ret := dp_minimumSubstringsInPartition(s, i+1, l, make([]int, 26), memo)
		if i-idx+1 == max_cnt*letter_cnt {
			res = min_int(res, 1+ret)
		}
	}
	memo[idx] = res
	return res
}

func MinimumSubstringsInPartition(s string) int {
	var l int = len(s)
	var cnt []int = make([]int, 26)
	var memo []int = make([]int, l)
	for i := 0; i < l; i++ {
		memo[i] = -1
	}
	return dp_minimumSubstringsInPartition(s, 0, l, cnt, memo)
}
