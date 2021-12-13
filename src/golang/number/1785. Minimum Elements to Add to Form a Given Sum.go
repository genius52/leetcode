package number

func minElements(nums []int, limit int, goal int) int {
	var sum int = 0
	for _, n := range nums {
		sum += n
	}
	target := goal - sum
	if target < 0 {
		target = -target
	}
	var res int = target / limit
	if target%limit != 0 {
		res++
	}
	return res
}
