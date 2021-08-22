package string_issue

func abs_int(n int)int{
	if n < 0{
		return -n
	}
	return n
}

func MinTimeToType(word string) int {
	var l int = len(word)
	var pre uint8 = 'a'
	var steps int = 0
	for i := 0;i < l;i++{
		cur := word[i]
		steps += min_int(abs_int(int(cur) - int(pre)),min_int(int(word[i]),int(pre)) + 26 - max_int(int(word[i]),int(pre)))
		steps++
		pre = word[i]
	}
	return steps
}