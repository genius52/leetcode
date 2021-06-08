package array

import "sort"

//Input: nums = [1,2,3,4,5,6,7,8]
//Output: true
//Explanation: We can split the array into [1,4,5,8] and [2,3,6,7], and both of them have an average of 4.5.
//average * l1

//sum / n = sum1 / k，可以变个形，sum * k / n = sum1
func dp_splitArraySameAverage(nums []int,l int,cur_sum int,target int,cur_pos int,rest_len int)bool{
	if cur_sum > target{
		return false
	}
	if cur_sum == target && rest_len == 0{
		return true
	}
	if cur_pos >= l{
		return false
	}
	if rest_len == 0{
		return false
	}
	if nums[cur_pos] > ((target - cur_sum)/rest_len){
		return false
	}
	for i := cur_pos;i < l - rest_len;i++{
		if i > cur_pos && nums[i] == nums[i - 1]{
			continue
		}
		res := dp_splitArraySameAverage(nums,l,cur_sum + nums[i],target,i + 1,rest_len - 1)
		if res{
			return true
		}
	}
	//?
	return false
	//skip cur
	//return dp_splitArraySameAverage(nums,l,cur_sum,target,cur_pos + 1,rest_len)
}

func SplitArraySameAverage(nums []int) bool {
	var sum int = 0
	sort.Ints(nums)
	var l int = len(nums)
	for i := 0;i < l;i++{
		sum += nums[i]
	}
	var  possible bool = false
	for i := 1; i <= l/2 && !possible;i++ {
		if sum * i % l == 0 {
			possible = true
		}
	}
	if (!possible){
		return false
	}
	for i := 1;i <= l/2;i++{
		if sum * i % l != 0{
			continue
		}
		target := sum * i / l
		if dp_splitArraySameAverage(nums,l,0,target,0,i){
			return true
		}
	}
	return false
}