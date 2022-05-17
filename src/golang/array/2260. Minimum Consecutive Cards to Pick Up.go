package array

func minimumCardPickup(cards []int) int {
	var l int = len(cards)
	var val_idx map[int]int = make(map[int]int)
	var res int = 2147483647
	for i := 0;i < l;i++{
		if pre,ok := val_idx[cards[i]];ok{
			res = min_int(res,i - pre + 1)
		}
		val_idx[cards[i]] = i
	}
	if res == 2147483647{
		return -1
	}
	return res
}