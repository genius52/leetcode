package number

//nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
//nums[i] - nums[j] == rev(nums[j]) - rev(nums[i])
//nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
func revert_num(n int) int {
	var res int = 0
	for n > 0 {
		mod := n % 10
		res *= 10
		res += mod
		n /= 10
	}
	return res
}

func countNicePairs(nums []int) int {
	var l int = len(nums)
	var record map[int]int = make(map[int]int)
	var MOD int = 1e9 + 7
	var res int = 0
	for i := 0; i < l; i++ {
		cur := nums[i] - revert_num(nums[i])
		if cnt, ok := record[cur]; ok {
			res += cnt
			res %= MOD
		}
		record[cur]++
	}
	return res
}
