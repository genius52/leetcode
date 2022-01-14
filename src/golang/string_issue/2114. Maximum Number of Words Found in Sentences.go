package string_issue

func mostWordsFound(sentences []string) int {
	var l int = len(sentences)
	var max_cnt int = 0
	for i := 0; i < l; i++ {
		var cnt int = 1
		for j := 0; j < len(sentences[i]); j++ {
			if sentences[i][j] == ' ' {
				cnt++
			}
		}
		if cnt > max_cnt {
			max_cnt = cnt
		}
	}
	return max_cnt
}
