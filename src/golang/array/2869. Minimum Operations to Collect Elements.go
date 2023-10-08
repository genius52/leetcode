package array

func minOperations2869(nums []int, k int) int {
	var l int = len(nums)
	var visited []bool = make([]bool, k+1)
	var cnt int = 0
	for i := l - 1; i >= 0; i-- {
		if nums[i] <= k && !visited[nums[i]] {
			visited[nums[i]] = true
			cnt++
			if cnt == k {
				return l - i
			}
		}
	}
	return l
}
