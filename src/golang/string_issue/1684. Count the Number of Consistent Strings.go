package string_issue

//Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
//Output: 2
//Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.
func countConsistentStrings(allowed string, words []string) int{
	var dict [26]bool
	for _,c := range allowed{
		dict[c - 'a'] = true
	}
	var l int = len(words)
	for _,w := range words{
		for _,c := range w{
			if !dict[c - 'a']{
				l--
				break
			}
		}
	}
	return l
}

func CountConsistentStrings(allowed string, words []string) int {
	var dict [26]bool
	for _,c := range allowed{
		dict[c - 'a'] = true
	}
	var res int = 0
	for _,w := range words{
		var require bool = true
		for _,c := range w{
			if !dict[c - 'a']{
				require = false
				break
			}
		}
		if require{
			res++
		}
	}
	return res
}