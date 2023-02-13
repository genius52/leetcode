package string_issue

func vowelStrings(words []string, queries [][]int) []int {
	var l1 int = len(words)
	var l2 int = len(queries)
	var prefix []int = make([]int,l1 + 1)
	var cnt int = 0
	for i := 0;i < l1;i++{
		var cur_len int = len(words[i])
		if (words[i][0] == 'a' || words[i][0] == 'e' || words[i][0] == 'i' || words[i][0] == 'o' || words[i][0] == 'u') &&
			(words[i][cur_len - 1] == 'a' || words[i][cur_len - 1] == 'e' || words[i][cur_len - 1] == 'i' || words[i][cur_len - 1] == 'o' || words[i][cur_len - 1] == 'u'){
			cnt++
		}
		prefix[i + 1] = cnt
	}
	var res []int = make([]int,l2)
	for i := 0;i < l2;i++{
		left := queries[i][0]
		right := queries[i][1]
		res[i] = prefix[right + 1] - prefix[left]
	}
	return res
}