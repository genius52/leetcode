package string_issue

func countWords(words1 []string, words2 []string) int {
	var record map[string]int = make(map[string]int)
	for _,w := range words1{
		record[w]++
	}
	for k,v := range record{
		if v > 1{
			delete(record,k)
		}
	}
	for _,w := range words2{
		if _,ok := record[w];ok{
			record[w]++
		}
	}
	var res int = 0
	for _,v := range record{
		if v == 2{
			res++
		}
	}
	return res
}