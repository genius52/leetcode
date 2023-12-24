package array

func CountSubarrays2962(nums []int, k int) int64 {
	var l1 int = len(nums)
	var max_num int = 0
	for i := 0; i < l1; i++ {
		if nums[i] > max_num {
			max_num = nums[i]
		}
	}
	var idx []int
	for i := 0; i < l1; i++ {
		if nums[i] == max_num {
			idx = append(idx, i)
		}
	}
	if len(idx) < k {
		return 0
	}
	var res int64 = 0
	var l2 int = len(idx)
	for i := 0; i+k <= l2; i++ {
		//left := idx[i]
		//right := idx[i+k]
		var left_dis int = 0
		var right_dis int = l1 - idx[i+k-1]
		if i == 0 {
			left_dis = idx[i] + 1
		} else {
			left_dis = idx[i] - idx[i-1]
		}
		//if i+k == l2 {
		//	right_dis = l1 - idx[i+k-1]
		//} else {
		//	right_dis = idx[i+k] - idx[i+k-1]
		//}
		res += int64(left_dis) * int64(right_dis)
	}
	return res
}
