package string_issue

func similarPairs(words []string) int {
	var l int = len(words)
	var record [][26]bool = make([][26]bool,l)
	for i := 0;i < l;i++{
		for _,c := range words[i]{
			record[i][c - 'a'] = true
		}
	}
	var res int = 0
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			var same bool = true
			for k := 0;k < 26;k++{
				if record[i][k] != record[j][k]{
					same = false
					break
				}
			}
			if same{
				res++
			}
		}
	}
	return res
}