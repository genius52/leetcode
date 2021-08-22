package number

func findGCD(nums []int) int {
	var l int = len(nums)
	var min_num int = 2147483647
	var max_num int = -2147483648
	for i := 0;i < l;i++{
		if nums[i] < min_num{
			min_num = nums[i]
		}
		if nums[i] > max_num{
			max_num = nums[i]
		}
	}
	return gcd(min_num,max_num)
}