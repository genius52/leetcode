package array

import (
	"sort"
)

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var record map[float32]int = make(map[float32]int)
	for i := 0;i < l/2;i++{
		var average float32 = (float32(nums[i]) + float32(nums[l - 1 - i]))/2
		record[average]++
	}
	return len(record)
}

const MIN = 0.000001

func DistinctAverages(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var res int = 0
	var data []float64
	for i := 0;i < l/2;i++{
		var average float64 = float64(nums[i]) + float64(nums[l - 1 - i])
		data = append(data,average)
	}
	var pre float64 = -1
	sort.Float64s(data)
	for i := 0;i < l/2;i++{
		if data[i] != pre{
			res++
		}
		pre = data[i]
	}
	return res
}