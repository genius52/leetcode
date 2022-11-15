package array

func SubarrayLCM(nums []int, k int) int {
	var l int = len(nums)
	var res int = 0
	for i := 0;i < l;i++{
		if nums[i] > k || k % nums[i] != 0{
			continue
		}
		var common int = nums[i]
		if common == k{
			res++
		}
		for j := i + 1;j < l;j++{
			if nums[j] > k || k % nums[j] != 0{
				break
			}
			cur := (nums[j] * common)/gcd(common,nums[j])
			if cur > k{
				break
			}
			common = cur
			if cur == k{
				res++
			}
		}
	}
	return res
}