package number

//func gcd(x, y int) int {
//	tmp := x % y
//	if tmp > 0 {
//		return gcd(y, tmp)
//	} else {
//		return y
//	}
//}

func dp_maxScore(nums []int, l int, status int, cnt int, memo map[int]int) int {
	var res int = 0
	if _, ok := memo[status]; ok {
		return memo[status]
	}
	for i := 0; i < l; i++ {
		if status|1<<i == status {
			continue
		}
		for j := 0; j < l; j++ {
			if i == j || status|1<<j == status {
				continue
			}
			next := status | 1<<i | 1<<j
			cur := cnt*gcd(nums[i], nums[j]) + dp_maxScore(nums, l, next, cnt+1, memo)
			res = max_int(res, cur)
		}
	}
	memo[status] = res
	return res
}

func MaxScore1799(nums []int) int {
	var l int = len(nums)
	var memo map[int]int = make(map[int]int)
	return dp_maxScore(nums, l, 0, 1, memo)
}
