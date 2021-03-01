package array

//Input: nums = [1,2,3,4,5,6,7], k = 3
//Output: [5,6,7,1,2,3,4]
//Explanation:
//rotate 1 steps to the right: [7,1,2,3,4,5,6]
//rotate 2 steps to the right: [6,7,1,2,3,4,5]
//rotate 3 steps to the right: [5,6,7,1,2,3,4]
func rotate189(nums []int, k int)  {
	var l int = len(nums)
	k = k % l
	if k == 0 || k == l{
		return
	}
	for i := 0;i < l/2;i++{
		nums[i],nums[l - 1 - i] = nums[l - 1 - i],nums[i]
	}
	for i := 0;i < k/2;i++{
		nums[i],nums[k - 1 - i] = nums[k - 1 - i],nums[i]
	}
	j := 0
	for i := k;i <= k + (l - 1-  k) / 2;i++{
		nums[i],nums[l - 1 - j] = nums[l - 1 - j],nums[i]
		j++
	}
}