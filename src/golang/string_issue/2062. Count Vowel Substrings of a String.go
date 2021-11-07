package string_issue


func countVowelSubstrings(word string) int {
	var res int
	var l int = len(word)
	for i := 0;i < l;i++{
		if !(word[i] == 'a' || word[i]  == 'e' ||word[i]  == 'i' ||word[i]  == 'o' ||word[i]  == 'u'){
			continue
		}
		mask := 0
		for j := i;j < l;j++{
			if word[j] == 'a'{
				mask = mask | 1
			}else if word[j] == 'e'{
				mask = mask | 1 << 1
			}else if word[j] == 'i'{
				mask = mask | 1 << 2
			}else if word[j] == 'o'{
				mask = mask | 1 << 3
			}else if word[j] == 'u'{
				mask = mask | 1 << 4
			}else{
				break
			}
			if mask == 1 << 5 - 1{
				res++
			}
		}
	}
	return res
}