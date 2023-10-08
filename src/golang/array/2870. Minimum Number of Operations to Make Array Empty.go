package array

func dfs_minOperations2870(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	var res int = -1
	if n == 1 {
		res = -1
	} else if n == 2 || n == 3 {
		res = 1
	} else {
		cur1 := dfs_minOperations2870(n-2, memo)
		cur2 := dfs_minOperations2870(n-3, memo)
		if cur1 != -1 && cur2 != -1 {
			res = 1 + min_int(cur1, cur2)
		} else if cur1 == -1 && cur2 != -1 {
			res = 1 + cur2
		} else if cur1 != -1 && cur2 == -1 {
			res = 1 + cur1
		}
	}
	memo[n] = res
	return res
}

func minOperations2870(nums []int) int {
	var num_cnt map[int]int = make(map[int]int)
	for _, n := range nums {
		num_cnt[n]++
	}
	var res int = 0
	var memo map[int]int = make(map[int]int)
	for _, cnt := range num_cnt {
		cur := dfs_minOperations2870(cnt, memo)
		if cur == -1 {
			return -1
		}
		res += cur
	}
	return res
}

func minOperations2870_2(nums []int) int {
	var num_cnt map[int]int = make(map[int]int)
	for _, n := range nums {
		num_cnt[n]++
	}
	var res int = 0
	for _, cnt := range num_cnt {
		if cnt == 1 {
			return -1
		}
		if cnt%3 == 0 {
			res += cnt / 3
		} else {
			res += cnt/3 + 1
		}
	}
	return res
}
