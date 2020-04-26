package number

import "sort"

//Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
//
//A solution set is:
//[
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
//]
func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	l := len(nums)
	for i := 0;i < l - 3;{
		if i != 0{
			if nums[i] == nums[i - 1]{
				i++
				continue
			}
		}
		for j := i + 1;j <= l - 3;{
			if(nums[i] + nums[j] + nums[j + 1] + nums[j + 2] < target){
				j++
				continue
			}else{
				if j + 3 < l{
					if nums[j + 2] == nums[j + 3]{
						j++
						continue
					}
				}
				second := i + 1
				third := j + 1
				for second  < third {
					sum := nums[i] + nums[second] + nums[third] + nums[j + 2]
					if sum == target{
						res = append(res,[]int{nums[i],nums[second], nums[third],nums[j + 2]})
						second++
						third--
						for second < l && nums[second] == nums[second - 1]{
							second++
						}
						for third >= 0 && nums[third] == nums[third + 1]{
							third--
						}
					}else if sum < target{
						second++
					}else{
						third--
					}
				}
				j++
			}
		}
		i++
	}
	return res
}