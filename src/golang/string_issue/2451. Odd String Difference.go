package string_issue

func OddString(words []string) string {
	var l int = len(words)
	var l2 int = len(words[0])
	for i := 1;i < l - 1;i++{
		for j := 0;j < l2 - 1;j++{
			if (words[i - 1][j + 1] - words[i - 1][j]) != (words[i][j + 1] - words[i][j]) ||
				(words[i][j + 1] - words[i][j]) != (words[i + 1][j + 1] - words[i + 1][j]){
				if (words[i - 1][j + 1] - words[i - 1][j]) == (words[i][j + 1] - words[i][j]){
					return words[i + 1]
				}else if (words[i - 1][j + 1] - words[i - 1][j]) == (words[i + 1][j + 1] - words[i + 1][j]){
					return words[i]
				}else{
					return words[i - 1]
				}
			}
		}
	}
	return ""
}