package array

import (
	"math"
	"sort"
)

func minMoves2(nums []int) int {
	var l int = len(nums)
	sort.Ints(nums)
	var target int = 0
	if (l | 1) == l{
		target = nums[l / 2]
	}else{
		target = (nums[l / 2] + nums[(l + 1)/2])/2
	}
	var cnt int = 0
	for i := 0;i < l;i++{
		cnt += int(math.Abs(float64(target - nums[i])))
	}
	return cnt
}