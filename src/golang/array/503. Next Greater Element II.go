package array

func NextGreaterElements(nums []int) []int {
	var l int = len(nums)
	if l == 0{
		return []int{}
	}
	nums = append(nums,nums...)
	var res []int = make([]int,l)
	res[0] = -1
	var stack []int
	stack = append(stack,0)
	for i := 1;i < l * 2;i++{
		if i < l{
			res[i % l] = -1
		}
		for len(stack) > 0 && nums[stack[len(stack) - 1]] < nums[i]{
			var index int
			index, stack = stack[len(stack)-1], stack[:len(stack)-1]
			res[index] = nums[i % l]
		}
		if i < l{
			stack = append(stack,i % l)
		}
	}
	return res
}