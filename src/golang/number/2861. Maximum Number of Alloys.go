package number

func MaxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	var res int = 0
	for i := 0; i < k; i++ {
		var low int = 0
		var high int = (stock[0] + budget/cost[0]) / composition[i][0]
		for low <= high {
			cnt := (low + high) / 2
			budget2 := budget
			var success bool = true
			for j := 0; j < len(composition[i]); j++ {
				if stock[j]/composition[i][j] < cnt {
					budget2 -= (composition[i][j]*cnt - stock[j]) * cost[j]
					if budget2 < 0 {
						success = false
						break
					}
				}
			}
			if success {
				low = cnt + 1
				res = max_int(res, cnt)
			} else {
				high = cnt - 1
			}
		}
	}
	return res
}
