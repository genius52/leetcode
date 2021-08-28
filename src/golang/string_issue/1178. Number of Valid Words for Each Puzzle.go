package string_issue

func findNumOfValidWords(words []string, puzzles []string) []int {
	var word_len int = len(words)
	var word_bitrecord []int = make([]int,word_len)
	var puzzle_len int = len(puzzles)
	//var puzzle_bitrecord []int = make([]int,puzzle_len)
	for i := 0;i < word_len;i++{
		var n int = 0
		for _,c := range words[i]{
			n = n | 1 << (c - 'a')
		}
		word_bitrecord[i] = n
	}
	var res []int = make([]int,puzzle_len)
	for i := 0;i < puzzle_len;i++{
		var n int = 0
		for _,c := range puzzles[i]{
			n = n | 1 << (c - 'a')
		}
		//puzzle_bitrecord[i] = n
		var cnt int = 0
		for j := 0;j < word_len;j++{
			if (n | word_bitrecord[j]) != n{
				continue
			}
			if (word_bitrecord[j] | (1 << (puzzles[i][0] - 'a'))) == word_bitrecord[j]{
				cnt++
			}
		}
		res[i] = cnt
	}
	return res
}