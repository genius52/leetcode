package array

//Input: weights = [3,2,2,4,1,4], D = 3
//Output: 6
//Explanation:
//A ship capacity of 6 is the minimum to ship all the packages in 3 days like this:
//1st day: 3, 2
//2nd day: 2, 4
//3rd day: 1, 4
func ShipWithinDays(weights []int, D int) int {
	var total int = 0
	var l int = len(weights)
	var max_weight int = 0
	for i := 0;i < l;i++{
		total += weights[i]
		if weights[i] > max_weight{
			max_weight = weights[i]
		}
	}
	var least_cap int = total / D
	if total % D != 0{
		least_cap++
	}
	if least_cap < max_weight{
		least_cap = max_weight
	}
	for true{
		var duration int = 1
		var cur_weight int = 0
		for i := 0;i < l;i++{
			if cur_weight + weights[i] <= least_cap{
				cur_weight += weights[i]
			}else{
				cur_weight = weights[i]
				duration++
			}
		}
		if duration <= D{
			break
		}
		least_cap++
	}
	return least_cap
}