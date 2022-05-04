package number

//Input: piles = [3,6,7,11], H = 8
//Output: 4
func MinEatingSpeed(piles []int, H int) int {
	var total int = 0
	var l int = len(piles)
	for i := 0;i < l;i++{
		total += piles[i]
	}
	least_speed := total / H
	if total % H != 0{
		least_speed++
	}
	for true{
		var cost_hours int = 0
		for i := 0;i < l;i++{
			cost_hours += piles[i]/least_speed
			if piles[i] % least_speed != 0{
				cost_hours++
			}
		}
		if cost_hours <= H{
			break
		}else{
			least_speed++
		}
	}
	return least_speed
}

func minEatingSpeed(piles []int, h int) int{
	var l int = len(piles)
	var left int = 1
	var right int = 0
	for i := 0;i < l;i++{
		right += piles[i]
	}
	for left < right{
		mid := (left + right)/2
		var hours int = 0
		for i := 0;i < l;i++{
			hours += piles[i]/ mid
			if piles[i] % mid != 0{
				hours++
			}
		}
		if hours > h{
			left = mid + 1
		}else{
			right = mid
		}
	}
	return left
}