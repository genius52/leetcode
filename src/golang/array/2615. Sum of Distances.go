package array

func Distance(nums []int) []int64 {
	var l int = len(nums)
	var num_idx map[int][]int = make(map[int][]int)
	for i := 0; i < l; i++ {
		num_idx[nums[i]] = append(num_idx[nums[i]], i)
	}
	var res []int64 = make([]int64, l)
	for _, idx := range num_idx {
		var l2 int = len(idx)
		if l2 > 1 {
			var prefix []int64 = make([]int64, l2+1)
			var suffix []int64 = make([]int64, l2+1)
			for j := 1; j < l2; j++ {
				prefix[j] = prefix[j-1] + int64((idx[j]-idx[j-1])*j)
			}
			for j := l2 - 2; j >= 0; j-- {
				suffix[j] = suffix[j+1] + int64((idx[j+1]-idx[j])*(l2-1-j))
			}
			for i := 0; i < l2; i++ {
				res[idx[i]] = prefix[i] + suffix[i]
			}
		}
	}
	return res
}
