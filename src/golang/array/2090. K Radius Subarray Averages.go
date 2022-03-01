package array

func getAverages(nums []int, k int) []int {
	if k == 0{
		return nums
	}
	var l int = len(nums)
	var prefix []int = make([]int,l + 1)
	for i := 0;i < l;i++{
		prefix[i + 1] = prefix[i] + nums[i]
	}
	var res []int = make([]int,l)
	for i := 0;i < min_int(k,l);i++{
		res[i] = -1
	}
	for i := l - 1;i > max_int(l - 1 - k,0);i--{
		res[i] = -1
	}
	for i := k;i <= l - 1 - k;i++{
		sum := prefix[i + k + 1] - prefix[i - k]
		res[i] = sum / (k * 2 + 1)
	}
	return res
}

//Input: nums = [7,4,3,9,1,8,5,2,6], k = 3
//Output: [-1,-1,-1,5,4,4,-1,-1,-1]
func GetAverages(nums []int, k int) []int{
	var l int = len(nums)
	var res []int = make([]int,l)
	var sum int = 0
	var cnt int = 1 + k * 2
	for i := 0;i < l;i++{
		sum += nums[i]
		res[i] = -1
		if i - k * 2 > 0{
			sum -= nums[i - k * 2 - 1]
		}
		if i - k * 2  >= 0{
			res[i - k] = sum/cnt
		}
	}
	return res
}