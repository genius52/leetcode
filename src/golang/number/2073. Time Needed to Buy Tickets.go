package number

func timeRequiredToBuy(tickets []int, k int) int {
	var l int = len(tickets)
	cnt := tickets[k]
	var res int = 0
	for i := 0;i < l;i++{
		if i <= k{
			if tickets[i] <= cnt{
				res += tickets[i]
			}else{
				res += cnt
			}
		}else{
			if tickets[i] < cnt{
				res += tickets[i]
			}else{
				res += cnt - 1
			}
		}
	}
	return res
}