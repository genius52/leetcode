package string_issue

//Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
//Output: [1,2]
//Explanation: On the first query only f("bbb") < f("aaaa"). On the second query both f("aaa") and f("aaaa") are both > f("cc").
func NumSmallerByFrequency(queries []string, words []string) []int {
	//brute force
	var l_word int = len(words)
	var words_cnt []int = make([]int,l_word)
	for i,word := range words{
		var record []int = make([]int,26)
		var min_char int32 = 'z'
		for _,c := range word{
			if(c <= min_char){
				min_char = c
			}
			record[c - 'a']++
		}
		words_cnt[i] = record[min_char - 'a']
	}
	var l int = len(queries)
	var res []int = make([]int,l)
	for i,q := range queries{
		var record []int = make([]int,26)
		var min_char int32 = 'z'
		for _,c := range q{
			if(c <= min_char){
				min_char = c
			}
			record[c - 'a']++
		}
		var bigger_cnt int = 0
		for _,cnt := range words_cnt{
			if (record[min_char - 'a'] < cnt){
				bigger_cnt++
			}
		}
		res[i] = bigger_cnt
	}
	return res
}