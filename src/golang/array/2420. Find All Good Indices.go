package array

func GoodIndices(nums []int, k int) []int {
	var l int = len(nums)
	var prefix_small []int = make([]int,l)
	prefix_small[0] = 1
	for i := 1;i < l;i++{
		if nums[i] <= nums[i - 1]{
			prefix_small[i] = prefix_small[i - 1] + 1
		}else{
			prefix_small[i] = 1
		}
	}
	var suffix_small []int = make([]int,l)
	suffix_small[l - 1] = 1
	for i := l - 2;i >= 0;i--{
		if nums[i] <= nums[i + 1]{
			suffix_small[i] = suffix_small[i + 1] + 1
		}else{
			suffix_small[i] = 1
		}
	}
	var res []int
	for i := k;i < l - k;i++{
		if prefix_small[i - 1] >= k && suffix_small[i + 1] >= k{
			res = append(res,i)
		}
	}
	return res
}