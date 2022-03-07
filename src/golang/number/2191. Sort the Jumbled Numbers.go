package number

import "sort"

func SortJumbled(mapping []int, nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		var n1 int = nums[i]
		var n2 int = nums[j]
		var res1 []int
		var res2 []int
		if n1 == 0{
			res1 = []int{mapping[n1]}
		}else{
			for n1 != 0{
				mod := n1 % 10
				res1 = append(res1,mapping[mod])
				n1 /= 10
			}
		}
		if n2 == 0{
			res2 = []int{mapping[n2]}
		}else{
			for n2 != 0{
				mod := n2 % 10
				res2 = append(res2,mapping[mod])
				n2 /= 10
			}
		}

		var num1 int = 0
		for k := len(res1) - 1;k >= 0;k--{
			num1 *= 10
			num1 += res1[k]
		}
		var num2 int = 0
		for k := len(res2) - 1;k >= 0;k--{
			num2 *= 10
			num2 += res2[k]
		}
		if num1 == num2{
			return i < j
		}else{
			return num1 < num2
		}
	})
	return nums
}