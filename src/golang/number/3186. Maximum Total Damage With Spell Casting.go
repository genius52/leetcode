package number

import "sort"

func dp_maximumTotalDamage(data [][2]int, l int, idx int, memo []int64) int64 {
	if idx >= l {
		return 0
	}
	if memo[idx] != -1 {
		return memo[idx]
	}
	var res int64 = 0
	//choose
	var next_idx int = idx + 1
	for next_idx < l && (data[next_idx][0]-data[idx][0]) <= 2 {
		next_idx++
	}
	var choose int64 = int64(data[idx][0])*int64(data[idx][1]) + dp_maximumTotalDamage(data, l, next_idx, memo)
	//skip
	var skip int64 = dp_maximumTotalDamage(data, l, idx+1, memo)
	res = max_int64(skip, choose)
	memo[idx] = res
	return res
}

func MaximumTotalDamage(power []int) int64 {
	var num_cnt map[int]int = make(map[int]int)
	for _, p := range power {
		num_cnt[p]++
	}
	var data [][2]int
	for p, cnt := range num_cnt {
		data = append(data, [2]int{p, cnt})
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
	var l int = len(data)
	var memo []int64 = make([]int64, l)
	for i := 0; i < l; i++ {
		memo[i] = -1
	}
	//var next []int = make([]int,l)
	return dp_maximumTotalDamage(data, len(data), 0, memo)
}
