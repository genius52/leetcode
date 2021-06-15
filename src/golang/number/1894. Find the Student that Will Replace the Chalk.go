package number

func chalkReplacer(chalk []int, k int) int {
	var n int = len(chalk)
	if n == 1{
		return 0
	}
	var sum int = 0
	for i := 0;i < n;i++{
		sum += chalk[i]
	}
	k = k % sum
	var cur int = 0
	for k >= chalk[cur]{
		k -= chalk[cur]
		cur = (cur + 1) % n
	}
	return cur
}