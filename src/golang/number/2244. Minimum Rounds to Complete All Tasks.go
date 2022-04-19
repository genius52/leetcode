package number

func dfs_minimumRounds(n int)int{
	if n % 3 == 0{
		return n / 3
	}
	return 1 + dfs_minimumRounds(n - 2)
}

func minimumRounds(tasks []int) int {
	var record map[int]int = make(map[int]int)
	for _,t := range tasks{
		record[t]++
	}
	var res int = 0
	for _,v := range record{
		if v == 1{
			return -1
		}
		res += dfs_minimumRounds(v)
	}
	return res
}