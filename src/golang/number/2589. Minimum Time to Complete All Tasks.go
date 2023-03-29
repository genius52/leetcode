package number

import "sort"

func findMinimumTime(tasks [][]int) int {
	var l int = len(tasks)
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	var res int = 0
	var task_cnt [2001]int
	for i := 0; i < l; i++ {
		start := tasks[i][0]
		end := tasks[i][1]
		duration := tasks[i][2]
		var current_cnt int = 0
		for j := start; j <= end; j++ {
			if task_cnt[j] != 0 {
				current_cnt++
			}
		}
		if current_cnt < duration {
			add_cnt := duration - current_cnt
			res += add_cnt
			for j := end; j >= start; j-- {
				if task_cnt[j] == 0 {
					task_cnt[j]++
					add_cnt--
					if add_cnt == 0 {
						break
					}
				}
			}
		}
	}
	return res
}
