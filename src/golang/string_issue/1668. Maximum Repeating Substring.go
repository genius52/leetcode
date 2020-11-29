package string_issue

func MaxRepeating(sequence string, word string) int {
	var l int = len(sequence)
	var word_len int = len(word)
	var res int = 0
	var visit int = 0
	for visit <= l - word_len{
		if sequence[visit] != word[0]{
			visit++
			continue
		}
		var i int = visit
		var cur_cnt int = 0
		for i < l{
			var j int = 0
			for i < l && j < word_len{
				if sequence[i] == word[j]{
					i++
					j++
					continue
				}else{
					break
				}
			}

			if j == word_len{
				cur_cnt++
				if cur_cnt > res{
					res = cur_cnt
				}
				visit = i
				j = 0
			}else{
				visit++
				break
			}
		}
	}
	return res
}