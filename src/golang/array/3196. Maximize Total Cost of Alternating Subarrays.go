package array

func dp_maximumTotalCost(nums []int, l int, idx int, isflip bool, memo [][2]int64) int64 {
	if idx >= l {
		return 0
	}
	if isflip {
		if memo[idx][1] != -(1 << 63) {
			return memo[idx][1]
		}
	} else {
		if memo[idx][0] != -(1 << 63) {
			return memo[idx][0]
		}
	}
	//continue
	var ret1 int64 = int64(nums[idx])
	if isflip {
		ret1 = -ret1
	}
	ret1 += dp_maximumTotalCost(nums, l, idx+1, !isflip, memo)
	//new subarray
	var ret2 int64 = dp_maximumTotalCost(nums, l, idx+1, false, memo)
	if isflip {
		memo[idx][1] = max_int64(ret1, ret2)
	} else {
		memo[idx][0] = max_int64(ret1, ret2)
	}
	return max_int64(ret1, ret2)
}

func MaximumTotalCost(nums []int) int64 {
	var l int = len(nums)
	var memo [][2]int64 = make([][2]int64, l)
	for i := 0; i < l; i++ {
		memo[i][0] = -(1 << 63)
		memo[i][1] = -(1 << 63)
	}
	return dp_maximumTotalCost(nums, l, 0, false, memo)
}
