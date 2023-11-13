package number

func abs_int64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func MinSum(nums1 []int, nums2 []int) int64 {
	var zero1 int64 = 0
	var zero2 int64 = 0
	var sum1 int64 = 0
	var sum2 int64 = 0
	for _, n := range nums1 {
		if n == 0 {
			zero1++
		}
		sum1 += int64(n)
	}
	for _, n := range nums2 {
		if n == 0 {
			zero2++
		}
		sum2 += int64(n)
	}
	if (zero1 == 0 && (sum1 < sum2 || sum2+zero2 > sum1)) || (zero2 == 0 && (sum2 < sum1 || sum1+zero1 > sum2)) {
		return -1
	}
	if sum1 < sum2 {
		//sum_diff := sum2 - sum1
		//zero_diff := zero1 - zero2
		if sum1+zero1 <= sum2+zero2 {
			return sum2 + zero2
		} else {
			return sum1 + zero1
		}
	} else if sum1 > sum2 {
		//sum_diff := sum1 - sum2
		//zero_diff := zero2 - zero1
		if sum2+zero2 <= sum1+zero1 {
			return sum1 + zero1
		} else {
			return sum2 + zero2
		}
	} else {
		return sum1 + max_int64(zero1, zero2)
	}
}
