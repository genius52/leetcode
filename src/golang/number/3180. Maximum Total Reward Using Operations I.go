package number

import "sort"

func dp_maxTotalReward(rewardValues []int, l int, idx int, sum int, memo map[int]int) int {
	if idx == l {
		return 0
	}
	if _, ok := memo[sum]; ok {
		return memo[sum]
	}
	if sum >= rewardValues[l-1] {
		return 0
	}
	for i := idx; i < l; i++ {
		if rewardValues[i] > sum {
			memo[sum] = max_int(memo[sum], rewardValues[i]+dp_maxTotalReward(rewardValues, l, i+1, sum+rewardValues[i], memo))
		}
	}
	return memo[sum]
}

func MaxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	var l int = len(rewardValues)
	var memo map[int]int = make(map[int]int)
	return dp_maxTotalReward(rewardValues, l, 0, 0, memo)
}
