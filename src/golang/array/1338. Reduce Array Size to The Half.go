package array

import (
	"math"
	"sort"
)

//Input: arr = [3,3,3,3,5,5,5,2,2,7]
//Output: 2
func minSetSize(arr []int) int {
	var record map[int]int = make(map[int]int)
	for _,n := range arr{
		if _,ok := record[n];ok{
			record[n]++
		}else{
			record[n] = 1
		}
	}
	var record_lengths []int
	for _, l := range record{
		record_lengths = append(record_lengths, l)
	}
	sort.Ints(record_lengths)
	target := int(math.Ceil(float64(len(arr))/2))
	total := 0
	for i := 0;i < len(record_lengths);i++{
		total += record_lengths[len(record_lengths) - i - 1]
		if total >= target{
			return i + 1
		}
	}
	return len(record_lengths)
}
