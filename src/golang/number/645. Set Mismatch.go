package number

//You have a set of integers s, which originally contains all the numbers from 1 to n.
//Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set, which results in repetition of one number and loss of another number.
//
//You are given an integer array nums representing the data status of this set after the error.
//
//Find the number that occurs twice and the number that is missing and return them in the form of an array.
//
//
//
//Example 1:
//
//Input: nums = [1,2,2,4]
//Output: [2,3]
func FindErrorNums(nums []int) []int {
	var l int = len(nums)
	var index int = 0
	for index < l{
		for nums[index] != (index + 1) && nums[nums[index] - 1] != nums[index]{
		//for nums[index] != (index + 1){
			nums[nums[index] - 1],nums[index] = nums[index],nums[nums[index] - 1]
		}
		index++
	}
	var res []int = make([]int,2)
	for i := 0;i < l;i++{
		if nums[i] != i + 1{
			res[0] = nums[i]
			res[1] = i + 1
			break
		}
	}
	return res
}