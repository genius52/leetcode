package string_issue

func numJewelsInStones(jewels string, stones string) int {
	var dict map[int32]bool = make(map[int32]bool)
	for _,j := range jewels{
		dict[j] = true
	}
	var res int = 0
	for _,s := range stones{
		if _,ok := dict[s];ok{
			res++
		}
	}
	return res
}