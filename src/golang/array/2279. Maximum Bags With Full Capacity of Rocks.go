package array

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	var l int = len(capacity)
	var res int = 0
	var diff []int
	for i := 0;i < l;i++{
		if capacity[i] == rocks[i]{
			res++
		}else{
			diff = append(diff,capacity[i] - rocks[i])
		}
	}
	sort.Ints(diff)
	for _,d := range diff{
		if additionalRocks < d{
			break
		}
		additionalRocks -= d
		res++
	}
	return res
}