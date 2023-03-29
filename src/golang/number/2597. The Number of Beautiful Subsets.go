package number

import "sort"

func dfs_beautifulSubsets(nums []int, l int, idx int, k int, num_cnt [1001]int) int {
	if idx == l {
		return 1
	}
	//skip current number
	var res int = dfs_beautifulSubsets(nums, l, idx+1, k, num_cnt)
	//choose current number
	if nums[idx] < k || nums[idx] >= k && num_cnt[nums[idx]-k] == 0 {
		num_cnt[nums[idx]]++
		res += dfs_beautifulSubsets(nums, l, idx+1, k, num_cnt)
		num_cnt[nums[idx]]--
	}
	return res
}

func BeautifulSubsets(nums []int, k int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var num_cnt [1001]int
	return dfs_beautifulSubsets(nums, l, 0, k, num_cnt) - 1
}
