package array


func FindDuplicates(nums []int) []int {
	var res []int
	for i := 0; i < len(nums);{
		if(nums[nums[i] - 1] != nums[i]){
			nums[nums[i] - 1],nums[i] = nums[i],nums[nums[i] - 1]
		} else{
			i++
		}
	}
	for i := 0;i < len(nums);i++{
		if(nums[i] - 1 != i){
			res = append(res, nums[i])
		}
	}
	return res
}