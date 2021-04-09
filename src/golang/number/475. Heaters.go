package number

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	res := 0
	for _,house_pos := range houses{
		min_diff := math.MaxInt32
		for _,heater_pos := range heaters{
			diff := math.Abs(float64(house_pos - heater_pos))
			if diff < float64(min_diff){
				min_diff = int(diff)
			}
		}
		if min_diff > res{
			res = min_diff
		}
	}
	return res
}


func FindRadius2(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var l1 int = len(houses)
	var l2 int = len(heaters)
	var dp []int = make([]int,l1)
	for i := 0;i < l1;i++{
		dp[i] = 2147483647
	}
	var house_index int = 0
	var heater_index int = 0
	for house_index < l1 && heater_index < l2{
		if houses[house_index] <= heaters[heater_index]{
			dp[house_index] = heaters[heater_index] - houses[house_index]
			house_index++
		}else{
			heater_index++
		}
	}
	house_index = l1 - 1
	heater_index = l2 - 1
	for house_index >= 0 && heater_index >= 0 {
		if houses[house_index] >= heaters[heater_index]{
			dp[house_index] = min_int(dp[house_index],houses[house_index] - heaters[heater_index])
			house_index--
		}else{
			heater_index--
		}
	}
	var max_dis int = 0
	for i := 0;i < l1;i++{
		max_dis = max_int(max_dis,dp[i])
	}
	return max_dis
}