package array

func dp_minIncrementOperations(nums []int, l int, idx int, last_idx int, k int, memo map[int]int64) int64 {
	if idx == l {
		return 0
	}
	var key int = 0
	if last_idx > 0 {
		key = idx*100000 + last_idx
	}
	if last_idx != -1 {
		if _, ok := memo[key]; ok {
			return memo[key]
		}
	}
	//nums[idx] >= k
	var ret int64 = 1<<62 - 1
	if nums[idx] >= k {
		ret = dp_minIncrementOperations(nums, l, idx+1, idx, k, memo)
	} else {
		if (idx - last_idx) <= 2 {
			ret = dp_minIncrementOperations(nums, l, idx+1, last_idx, k, memo)
		}
		ret = min_int64_number(ret, int64(k-nums[idx])+dp_minIncrementOperations(nums, l, idx+1, idx, k, memo))
	}
	if last_idx >= 0 {
		memo[key] = ret
	}
	return ret
}

func minIncrementOperations(nums []int, k int) int64 {
	var l int = len(nums)
	var memo map[int]int64 = make(map[int]int64)
	return dp_minIncrementOperations(nums, l, 0, -1, k, memo)
}
