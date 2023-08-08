package array

func dfs_canSplitArray(nums []int, prefix_sum []int, l int, m int, left int, right int, visited [][]bool) bool {
	if (left + 1) == right {
		return true
	}
	if visited[left][right] {
		return false
	}
	var res bool = false
	if (prefix_sum[right+1] - prefix_sum[left+1]) >= m { //cut left side
		res = res || dfs_canSplitArray(nums, prefix_sum, l, m, left+1, right, visited)
		if res {
			return true
		}
	}
	if (prefix_sum[right] - prefix_sum[left]) >= m { //cut right side
		res = res || dfs_canSplitArray(nums, prefix_sum, l, m, left, right-1, visited)
		if res {
			return true
		}
	}
	visited[left][right] = true
	return res
}

func CanSplitArray(nums []int, m int) bool {
	var l int = len(nums)
	if l <= 2 {
		return true
	}
	var prefix_sum []int = make([]int, l+1)
	for i := 0; i < l; i++ {
		prefix_sum[i+1] = prefix_sum[i] + nums[i]
	}
	var visited [][]bool = make([][]bool, l+1)
	for i := 0; i <= l; i++ {
		visited[i] = make([]bool, l+1)
	}
	return dfs_canSplitArray(nums, prefix_sum, l, m, 0, l-1, visited)
}
