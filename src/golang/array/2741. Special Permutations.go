package array

func dfs_specialPerm(nums []int, l int, pre_idx int, status int, cnt int, visited []bool, memo [][]int) int {
	if cnt == l {
		return 1
	}
	if pre_idx != -1 && memo[pre_idx][status] != -1 {
		return memo[pre_idx][status]
	}
	var res int = 0
	var MOD int = 1e9 + 7
	for i := 0; i < l; i++ {
		if visited[i] {
			continue
		}
		if pre_idx == -1 || (nums[i]%nums[pre_idx] == 0 || nums[pre_idx]%nums[i] == 0) {
			visited[i] = true
			res += dfs_specialPerm(nums, l, i, status|1<<i, cnt+1, visited, memo)
			res %= MOD
			visited[i] = false
		}
	}
	if pre_idx != -1 {
		memo[pre_idx][status] = res
	}
	return res
}

func SpecialPerm(nums []int) int {
	var l int = len(nums)
	var max_status int = 1 << l
	var memo [][]int = make([][]int, l)
	for i := 0; i < l; i++ {
		memo[i] = make([]int, max_status+1)
		for j := 0; j <= max_status; j++ {
			memo[i][j] = -1
		}
	}
	var visited []bool = make([]bool, l)
	return dfs_specialPerm(nums, l, -1, 0, 0, visited, memo)
}
