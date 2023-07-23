package array

func dfs_maxScore(nums []int, l int, idx int, x int, memo []int64) int64 {
	if idx == l {
		return 0
	}
	if memo[idx] != -1 {
		return memo[idx]
	}
	var res int64 = 0
	var find_same bool = false
	var find_diff bool = false
	for i := idx + 1; i < l; i++ {
		if ((nums[idx] + nums[i]) | 1) != (nums[idx] + nums[i]) { //same: odd odd / even even
			if find_same {
				continue
			}
			find_same = true
			cur := int64(nums[i]) + dfs_maxScore(nums, l, i, x, memo)
			if cur > res {
				res = cur
			}
		} else {
			find_diff = true
			cur := int64(nums[i]-x) + dfs_maxScore(nums, l, i, x, memo)
			if cur > res {
				res = cur
			}
		}
		if find_same && find_diff {
			break
		}
	}
	memo[idx] = res
	return res
}

func maxScore(nums []int, x int) int64 {
	var l int = len(nums)
	var memo []int64 = make([]int64, l)
	for i := 0; i < l; i++ {
		memo[i] = -1
	}
	return int64(nums[0]) + dfs_maxScore(nums, l, 0, x, memo)
}

// 2, 3, 6, 1, 9, 2
func MaxScore2(nums []int, x int) int64 {
	var l int = len(nums)
	var cur_odd_max int64 = -1
	var cur_even_max int64 = -1
	if (nums[0] | 1) == nums[0] {
		cur_odd_max = int64(nums[0])
	} else {
		cur_even_max = int64(nums[0])
	}
	var res int64 = int64(nums[0])
	for i := 1; i < l; i++ {
		if (nums[i] | 1) == nums[i] { //当前是奇数, even -> odd,odd -> odd
			if cur_even_max != -1 { //2,3
				if cur_odd_max == -1 {
					cur_odd_max = int64(nums[i]-x) + cur_even_max
				} else {
					cur_odd_max = max_int64(int64(nums[i]-x)+cur_even_max, int64(nums[i])+cur_odd_max)
				}
			} else {
				cur_odd_max += int64(nums[i])
			}
			res = max_int64(res, cur_odd_max)
		} else { //当前是偶数, even -> even, odd -> even
			if cur_odd_max != -1 {
				if cur_even_max == -1 {
					cur_even_max = int64(nums[i]-x) + cur_odd_max
				} else {
					cur_even_max = max_int64(int64(nums[i]-x)+cur_odd_max, int64(nums[i])+cur_even_max)
				}
			} else {
				cur_even_max += int64(nums[i])
			}
			res = max_int64(res, cur_even_max)
		}
	}
	return res
}
