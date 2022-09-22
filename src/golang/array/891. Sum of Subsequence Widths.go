package array

import (
	"math"
	"sort"
)

//Input: nums = [2,1,3]
//Output: 6
//Explanation:
//Subsequences are [1], [2], [3], [2,1], [2,3], [1,3], [2,1,3].
//The corresponding widths are 0, 0, 0, 1, 1, 2, 2.
//The sum of these widths is 6.
func SumSubseqWidths(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var res int64 = 0
	for i := 0;i < l;i++{
		//nums[i] as the biggest number
		res = (res + int64(nums[i]) * (int64(math.Pow(2,float64(i))) % 1000000007) - int64(nums[i]) * (int64(math.Pow(2,float64(l - i - 1))) % 1000000007)) % 1000000007
	}
	return int(res + 1000000007) % 1000000007
}

func sumSubseqWidths(nums []int) int{
	var l int = len(nums)
	sort.Ints(nums)
	var MOD int64 = 1e9 + 7
	var res int64 = 0
	var mutiple []int64 = make([]int64,l + 1)
	mutiple[0] = 1
	for i := 1;i <= l;i++{
		mutiple[i] = (mutiple[i - 1] << 1) % MOD
	}
	for i := 0;i < l;i++{
		smallest_cnt := l - 1 - i
		biggest_cnt := i
		res = (res + (mutiple[biggest_cnt] - mutiple[smallest_cnt]) * int64(nums[i])) % MOD
	}
	return int((res + MOD) % MOD)
}