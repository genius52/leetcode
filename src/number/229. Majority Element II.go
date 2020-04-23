package number

import (
	"math"
	"sort"
)

//Input: [1,1,1,3,3,2,2,2]
//Output: [1,2]
func MajorityElement(nums []int) []int {
	sort.Ints(nums)
	var res []int
	l := len(nums)
	step := int(math.Floor(float64(l) / 3))
	for i := 0;i < l;{
		if i + step >= l{
			break
		}
		if nums[i] == nums[i + step]{
			res = append(res,nums[i])
			i = i + step + 1
			for i < l && nums[i] == nums[i - 1]{
				i++
			}
		}else{
			i++
		}
	}
	return res
}
