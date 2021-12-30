package number

func maximumPopulation(logs [][]int) int {
	var record [101]int
	var early_year int = 100
	var last_year int = 0
	for _, log := range logs {
		birth := log[0] - 1950
		die := log[1] - 1950
		record[birth]++
		record[die]--
		if birth < early_year {
			early_year = birth
		}
		if die > last_year {
			last_year = die
		}
	}
	var max_cnt int = 0
	var max_cnt_year int = 0
	var cnt int = 0
	for i := early_year; i <= last_year; i++ {
		cnt += record[i]
		if cnt > max_cnt {
			max_cnt = cnt
			max_cnt_year = i
		}
	}
	return max_cnt_year + 1950
}
