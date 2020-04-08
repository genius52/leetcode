package number
//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//Example:
//
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

//brute force solution O(N * N)
func twoSum(nums []int, target int) []int {
	l := len(nums)
	var res []int
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			if nums[i] + nums[j] == target{
				res = append(res,i)
				res = append(res,j)
				return res
			}
		}
	}
	return res
}

//use Map store the previous num.  O(N * logN)
func TwoSum2(nums []int, target int) []int {
	l := len(nums)
	var res []int = make([]int,2)
	var record map[int]int = make(map[int]int)
	for i := 0;i < l;i++{
		diff := target - nums[i]
		if _,ok := record[diff];ok{
			res[0] = record[diff]
			res[1] = i
			return res
		}else{
			record[nums[i]] = i
		}
	}
	return res
}
