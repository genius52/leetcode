package array

func FindKDistantIndices(nums []int, key int, k int) []int {
	var res []int
	var l int = len(nums)
	var pre_idx int = -1
	for i := 0;i < l;i++{
		if nums[i] == key{
			var low int = i - k
			if low < 0{
				low = 0
			}
			if low < pre_idx{
				low = pre_idx
			}
			var high int = i + k
			if high >= l{
				high = l - 1
			}
			for j := low;j <= high;j++{
				res = append(res,j)
			}
			if high == l - 1{
				break
			}
			pre_idx = high + 1
		}
	}
	return res
}