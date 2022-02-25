package number

func timeRequiredToBuy2(tickets []int, k int) int{
	target := tickets[k]
	var l int = len(tickets)
	var res int = 0
	for i := 0;i < l;i++{
		if tickets[i] < target{
			res += tickets[i]
		}else{
			if i <= k{
				res += target
			}else {
				res += target - 1
			}
		}
	}
	return res
}

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