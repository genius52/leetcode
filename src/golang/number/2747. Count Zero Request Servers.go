package number

import "sort"

func CountServers(n int, logs [][]int, x int, queries []int) []int {
	var q_len int = len(queries)
	var querytime_idx [][2]int = make([][2]int, q_len)
	for i := 0; i < q_len; i++ {
		querytime_idx[i] = [2]int{queries[i], i}
	}
	sort.Slice(querytime_idx, func(i, j int) bool {
		return querytime_idx[i][0] < querytime_idx[j][0]
	})
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})
	var log_len int = len(logs)
	var res []int = make([]int, q_len)
	var left int = 0
	var right int = 0
	var machine_cnt map[int]int = make(map[int]int)
	for i := 0; i < q_len; i++ {
		start := querytime_idx[i][0] - x
		end := querytime_idx[i][0]
		if right < log_len && logs[right][1] < start {
			machine_cnt = make(map[int]int)
			left = right
			for left < log_len && logs[left][1] < start {
				left++
			}
			right = left
			for right < log_len && logs[right][1] <= end {
				machine_cnt[logs[right][0]]++
				right++
			}
		} else {
			for left < log_len && logs[left][1] < start {
				machine_cnt[logs[left][0]]--
				if machine_cnt[logs[left][0]] == 0 {
					delete(machine_cnt, logs[left][0])
				}
				left++
			}
			for right < log_len && logs[right][1] <= end {
				machine_cnt[logs[right][0]]++
				right++
			}
		}
		res[querytime_idx[i][1]] = n - len(machine_cnt)
	}
	return res
}
