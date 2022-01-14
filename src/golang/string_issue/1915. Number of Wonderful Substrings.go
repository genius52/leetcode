package string_issue

func WonderfulSubstrings(word string) int64 {
	var l int = len(word)
	var record []int64 = make([]int64, 1<<10) //record[i]前缀和为i的个数
	var res int64 = 0
	var pre int = 0
	record[0] = 1
	for i := 0; i < l; i++ {
		pre ^= 1 << (word[i] - 'a')
		res += record[pre] // 所有字母均出现偶数次
		for j := 0; j < 10; j++ {
			mask := 1 << j
			res += record[pre^mask]
		}
		record[pre]++
	}
	return res
}
