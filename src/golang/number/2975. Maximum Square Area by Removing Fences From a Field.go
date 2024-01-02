package number

import "sort"

// m:横线
// 年:竖线
func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	var MOD int64 = 1e9 + 7
	var res int = -1
	hFences = append(hFences, 1)
	hFences = append(hFences, m)
	vFences = append(vFences, 1)
	vFences = append(vFences, n)
	sort.Ints(hFences)
	sort.Ints(vFences)
	var h_record map[int]bool = make(map[int]bool)
	var v_record map[int]bool = make(map[int]bool)
	var l1 int = len(hFences)
	var l2 int = len(vFences)
	for i := 0; i < l1; i++ {
		for j := i + 1; j < l1; j++ {
			h_record[hFences[j]-hFences[i]] = true
		}
	}
	for i := 0; i < l2; i++ {
		for j := i + 1; j < l2; j++ {
			v_record[vFences[j]-vFences[i]] = true
		}
	}
	for dis, _ := range h_record {
		if _, ok := v_record[dis]; ok {
			if dis > res {
				res = dis
			}
		}
	}
	if res != -1 {
		return int(int64(res) * int64(res) % MOD)
	}
	return res
}
