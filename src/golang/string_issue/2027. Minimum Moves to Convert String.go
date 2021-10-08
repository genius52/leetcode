package string_issue

func minimumMoves(s string) int {
	var l int = len(s)
	var res int = 0
	for i := 0;i < l;{
		if s[i] == 'X'{
			i += 3
			res++
		}else{
			i++
		}
	}
	return res
}