package string_issue

//Input: sentence = "i love eating burger", searchWord = "burg"
//Output: 4
func IsPrefixOfWord(sentence string, searchWord string) int {
	var l1 int = len(sentence)
	var l2 int = len(searchWord)
	var left int = 0
	var res int= -1
	var idx int = 1
	for left < l1{
		var visit int = 0
		for left < l1 && visit < l2 && sentence[left] == searchWord[visit]{
			left++
			visit++
		}
		if visit == l2{
			return idx
		}
		for left < l1 && sentence[left] != ' '{
			left++
		}
		idx++
		left++
	}
	return res
}