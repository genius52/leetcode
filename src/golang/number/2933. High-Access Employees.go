package number

import (
	"sort"
	"strconv"
)

func FindHighAccessEmployees(access_times [][]string) []string {
	var res []string
	var record map[string][]int = make(map[string][]int)
	for _, at := range access_times {
		hours, _ := strconv.Atoi(at[1][:2])
		minutes, _ := strconv.Atoi(at[1][2:])
		record[at[0]] = append(record[at[0]], hours*60+minutes)
	}
	for k, v := range record {
		if len(v) < 3 {
			continue
		}
		sort.Ints(v)
		left := 0
		right := 2
		for right < len(v) {
			if v[right]-v[left] < 60 {
				res = append(res, k)
				break
			}
			left++
			right++
		}
	}
	return res
}
