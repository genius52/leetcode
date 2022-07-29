package number

import "sort"

func minOperations2344(nums []int, numsDivide []int) int {
	sort.Ints(nums)
	var l1 int = len(nums)
	var l2 int = len(numsDivide)
	var c int = numsDivide[0]
	for i := 1;i < l2;i++{
		c = gcd(c,numsDivide[i])
		if c == 1{
			break
		}
	}
	var res int = 0
	var i int = 0
	for i < l1{
		if c >= nums[i] && c % nums[i] == 0{
			break
		}
		i++
	}

	return res
}