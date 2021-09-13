package number

func findNumbers(nums []int) int {
	var res int = 0
	for i := 0;i < len(nums);i++{
		var even bool = true
		for nums[i] > 0{
			nums[i] /= 10
			even = !even
		}
		if even{
			res++
		}
	}
	return res
}
