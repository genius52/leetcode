package string_issue


func countVowelSubstrings(word string) int {
	var res int
	var l int = len(word)
	var require int = 1  | 1 << ('e' - 'a') | 1 << ('i' - 'a') | 1 << ('o' - 'a') | 1 << ('u' - 'a')
	for i := 0;i < l;i++{
		if !(word[i] == 'a' || word[i]  == 'e' ||word[i]  == 'i' ||word[i]  == 'o' ||word[i]  == 'u'){
			continue
		}
		mask := 0
		for j := i;j < l;j++{
			if word[j] == 'a' || word[j] == 'e' || word[j] == 'i' || word[j] == 'o' || word[j] == 'u'{
				mask = mask | 1 << (word[j] - 'a')
				if mask == require{
					res++
				}
			}else{
				break
			}
		}
	}
	return res
}