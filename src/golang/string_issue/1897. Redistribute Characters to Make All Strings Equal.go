package string_issue

func MakeEqual(words []string) bool {
	var l int = len(words)
	var record [26]int
	for i := 0;i < l;i++{
		var cur_len int = len(words[i])
		for j := 0;j < cur_len;j++{
			record[words[i][j] - 'a']++
		}
	}
	for i := 0;i < 26;i++{
		if record[i] == 0{
			continue
		}
		if (record[i] % l) != 0{
			return false
		}
	}
	return true
}