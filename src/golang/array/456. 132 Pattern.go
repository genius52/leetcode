package array

import "container/list"

//Input: nums = [3,1,4,2]
//Output: true
//Explanation: There is a 132 pattern in the sequence: [1, 4, 2].
func find132pattern(nums []int) bool {
	var l int = len(nums)
	if l < 3{
		return false
	}
	var dp_min []int = make([]int,l)
	dp_min[0] = nums[0]
	for i := 1;i < l;i++{
		dp_min[i] = min_int(nums[i],dp_min[i - 1])
		if nums[i] == dp_min[i]{
			continue
		}
		for j := i - 1;j > 0;j--{
			if dp_min[j - 1] >= nums[i]{
				break
			}
			if nums[j] > nums[i] && nums[j] > dp_min[j - 1] && dp_min[j - 1] < nums[i]{
				return true
			}
		}
	}
	return false
}

func Find132pattern(nums []int) bool {
	var l int = len(nums)
	if l < 3 {
		return false
	}
	//var left_min []int = make([]int,l + 1)
	//for i := 0;i < l;i++{
	//	left_min[i + 1] = min_int(left_min[i],nums[i])
	//}
	var q list.List // increasing stack,keep min number and max number
	//var pre_max_val int = 2147483647
	var second_max int = -2147483648
	for i := l - 1;i >= 0;i--{
		if nums[i] < second_max{
			return true
		}
		for q.Len() != 0 && nums[i] > q.Back().Value.(int){
			second_max = q.Back().Value.(int)
			q.Remove(q.Back())
		}
		q.PushBack(nums[i])
	}
	return false
}