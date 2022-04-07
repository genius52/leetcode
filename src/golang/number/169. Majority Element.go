package number

import "sort"

//Input: nums = [2,2,1,1,1,2,2]
//Output: 2
func MajorityElement1(nums []int) int {
	var l int = len(nums)
	if l == 1{
		return nums[0]
	}
	var res int32 = 0
	for i := 0;i < 32;i++{
		var mask int = 1 << i
		var cnt int = 0
		for _,n := range nums{
			if (n & mask) != 0{
				cnt++
				if cnt > l/2{
					break
				}
			}
		}
		if cnt > l/2{
			res |= int32(mask)
		}
	}
	return int(res)
}

func majorityElement(nums []int) int{
	sort.Ints(nums)
	return nums[len(nums)/2]
}