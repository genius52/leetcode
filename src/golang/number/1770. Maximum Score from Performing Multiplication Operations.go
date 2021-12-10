package number

//1 <= m <= 103
//m <= n <= 105
func dp_maximumScore(nums []int, l1 int, left int, right int, multipliers []int, l2 int, pos2 int, memo map[int]int) int {
	if left > right || pos2 >= l2 {
		return 0
	}
	key := left + right*100000
	if _, ok := memo[key]; ok {
		return memo[key]
	}
	var res int = -2147483648
	res = max_int(multipliers[pos2]*nums[left]+dp_maximumScore(nums, l1, left+1, right, multipliers, l2, pos2+1, memo),
		multipliers[pos2]*nums[right]+dp_maximumScore(nums, l1, left, right-1, multipliers, l2, pos2+1, memo))
	memo[key] = res
	return res
}

func maximumScore1770(nums []int, multipliers []int) int {
	var l1 int = len(nums)
	var l2 int = len(multipliers)
	var memo map[int]int = make(map[int]int)
	return dp_maximumScore(nums, l1, 0, l1-1, multipliers, l2, 0, memo)
}
