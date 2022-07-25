package number

func bestHand(ranks []int, suits []byte) string {
	var card_num [14]int
	var card_type [4]int
	var find_three bool = false
	var find_pair bool = false
	for idx,num := range ranks{
		card_num[num]++
		card_type[suits[idx] - 'a']++
		if card_num[num] >= 3{
			find_three = true
		}
		if card_num[num] == 2{
			find_pair = true
		}
	}
	for i := 0;i < 4;i++{
		if card_type[i] == 5{
			return "Flush"
		}
	}
	if find_three{
		return "Three of a Kind"
	}
	if find_pair{
		return "Pair"
	}
	return "High Card"
}