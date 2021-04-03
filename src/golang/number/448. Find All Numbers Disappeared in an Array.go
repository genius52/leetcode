package number

//Input: nums = [4,3,2,7,8,2,3,1]
//Output: [5,6]
func FindDisappearedNumbers(nums []int) []int {
	var l int = len(nums)
	var i int = 0
	for i < l{
		if nums[nums[i] - 1] != nums[i]{
			nums[nums[i] - 1],nums[i] = nums[i],nums[nums[i] - 1]
		}else{
			i++
		}
	}
	var res []int
	for i := 0;i < l;i++{
		if nums[i] != i + 1{
			res = append(res,i + 1)
		}
	}
	return res
}