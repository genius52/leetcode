package array

import "sort"

func MaxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	var record map[int]int = make(map[int]int)
	var l int = len(difficulty)
	var low int = 100000
	var high int = 1
	for i := 0;i < l;i++{
		if pre_profit,ok := record[difficulty[i]];ok{
			if pre_profit < profit[i]{
				record[difficulty[i]] = profit[i]
			}
		}else{
			record[difficulty[i]] = profit[i]
		}

		if difficulty[i] < low{
			low = difficulty[i]
		}
		if difficulty[i] > high{
			high = difficulty[i]
		}
	}
	var dp []int = make([]int,100001)
	var max_profit int = 0
	for i := low;i <= 100000;i++{
		if profit,ok := record[i];ok{
			if profit > max_profit{
				max_profit = profit
			}
			dp[i] = max_profit
		}else{
			dp[i] = max_profit
		}
	}

	var res int = 0
	var l2 int = len(worker)
	for i := 0;i < l2;i++{
		res += dp[worker[i]]
	}
	return res
}

//Two pointer solution
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	var l1 int = len(difficulty)
	var record [][2]int = make([][2]int,l1)
	for i := 0;i < l1;i++{
		record[i] = [2]int{difficulty[i],profit[i]}
	}
	sort.Slice(record, func(i, j int) bool {
		return record[i][0] < record[j][0]
	})
	sort.Ints(worker)
	var res int = 0
	var job_idx int = 0
	var max_profit int = 0
	for i := 0;i < len(worker);i++{
		for job_idx < l1 && worker[i] >= record[job_idx][0]{
			if record[job_idx][1] > max_profit{
				max_profit = record[job_idx][1]
			}
			job_idx++
		}
		res += max_profit
	}
	return res
}