package number

func dp_earliestAndLatest(status int, n int, round int, left int, right int, firstPlayer int, secondPlayer int, min_round *int, max_round *int, memo map[int]bool) {
	if status == 0 {
		return
	}
	for left < right && status|(1<<left) != status {
		left++
	}
	for left < right && status|(1<<right) != status {
		right--
	}
	if left < right {
		if (left+1) == firstPlayer && (right+1) == secondPlayer {
			*min_round = min_int(*min_round, round)
			*max_round = max_int(*max_round, round)
			return
		}

		if (left + 1) == firstPlayer {
			//status = status ^ (1 << right)
			dp_earliestAndLatest(status^(1<<right), n, round, left+1, right-1, firstPlayer, secondPlayer, min_round, max_round, memo)
		} else if (right + 1) == secondPlayer {
			//status = status ^ (1 << left)
			dp_earliestAndLatest(status^(1<<left), n, round, left+1, right-1, firstPlayer, secondPlayer, min_round, max_round, memo)
		} else {
			dp_earliestAndLatest(status^(1<<right), n, round, left+1, right-1, firstPlayer, secondPlayer, min_round, max_round, memo)
			dp_earliestAndLatest(status^(1<<left), n, round, left+1, right-1, firstPlayer, secondPlayer, min_round, max_round, memo)
		}
	} else {
		if _, ok := memo[status]; ok {
			return
		}
		memo[status] = true
		dp_earliestAndLatest(status, n, round+1, 0, n-1, firstPlayer, secondPlayer, min_round, max_round, memo)
	}
}

func EarliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	var memo map[int]bool = make(map[int]bool)
	var min_round int = 2147483647
	var max_round int = 0
	dp_earliestAndLatest((1<<n)-1, n, 1, 0, n-1, firstPlayer, secondPlayer, &min_round, &max_round, memo)
	return []int{min_round, max_round}
}
