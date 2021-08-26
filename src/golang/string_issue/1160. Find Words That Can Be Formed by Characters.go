package string_issue

func countCharacters(words []string, chars string) int {
	var dict [26]int
	var used_len int = len(chars)
	for _,c := range chars{
		dict[c - 'a']++
	}
	var res int = 0
	for _,word := range words{
		var l2 int = len(word)
		if l2 > used_len{
			continue
		}
		var cnt [26]int
		var meet bool = true
		for i := 0;i < l2;i++{
			c := word[i] - 'a'
			cnt[c]++
			if cnt[c] > dict[c]{
				meet = false
				break
			}
		}
		if meet{
			res += l2
		}
	}
	return res
}